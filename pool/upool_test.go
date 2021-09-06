package pool

import (
	"math/rand"
	"reflect"
	"testing"
	"time"
	"unsafe"
)

// TODO: use code generation like c++ template class
type UnsafePool struct {
	New func() *Test
	//sz       uintptr
	//lastptr  uintptr
	//ptrs     map[uintptr]struct{}
	freeptrs []uintptr
}

func NewUnsafePool(new func() *Test) *UnsafePool {
	return &UnsafePool{
		New: new,
		//ptrs:     make(map[uintptr]struct{}),
		freeptrs: make([]uintptr, 0),
	}
}

func (p *UnsafePool) noescape(up unsafe.Pointer) unsafe.Pointer {
	return unsafe.Pointer(uintptr(up) ^ 0)
}

func (p *UnsafePool) uintptr2t(ptr uintptr) *Test {
	return (*Test)(unsafe.Pointer(ptr))
}

func (p *UnsafePool) t2uintptr(v *Test) uintptr {
	return uintptr(unsafe.Pointer(v))
}

func (p *UnsafePool) Get() *Test {
	var st *Test
	if len(p.freeptrs) == 0 {
		return p.uintptr2t(uintptr(p.noescape(unsafe.Pointer(p.New()))))
	} else {
		ptr := p.freeptrs[len(p.freeptrs)-1]
		sh := (*reflect.SliceHeader)(unsafe.Pointer(&p.freeptrs))
		sh.Len--
		// or you can use -> p.freeptrs = p.freeptrs[:len(p.freeptrs)-1]
		st = p.uintptr2t(ptr)
	}
	return st
}

//func (p *UnsafePool) _Get() *Test {
// var st *Test
// if len(p.freeptrs) == 0 {
//  st = p.New()
//  //p.sz = unsafe.Sizeof(*st)
//  //ptr := p.t2uintptr(st)
//  //p.ptrs[ptr] = struct{}{}
//  //p.lastptr = ptr
// } else {
//  if len(p.freeptrs) != 0 {
//   ptr := p.freeptrs[len(p.freeptrs)-1]
//   p.freeptrs = p.freeptrs[:len(p.freeptrs)-1]
//   st = p.uintptr2t(ptr)
//  } else {
//   //nextPtr := unsafe.Add(unsafe.Pointer(p.lastptr), p.sz)
//   //fmt.Println(uintptr(nextPtr))
//   st = p.New()
//   //p.write2unsafePointer(st, nextPtr)
//   //next := uintptr(unsafe.Pointer(st))
//   //fmt.Println(next)
//   //p.ptrs[next] = struct{}{}
//   //p.lastptr = next
//   //st = p.uintptr2t(next)
//   fmt.Println(st)
//  }
// }
// return st
//}

func (p *UnsafePool) Return(st *Test) {
	p.freeptrs = append(p.freeptrs, p.t2uintptr(st))
}

var x *Test

//goos: linux
//goarch: amd64
//Benchmark_PoolGetOnly              10000           3267820 ns/op         2240008 B/op      20000 allocs/op
//Benchmark_PoolGetReturn            10000           1977241 ns/op         1120046 B/op      10000 allocs/op
func Benchmark_PoolGetOnly(b *testing.B) {
	f := func() *Test {
		te := Test{}
		te.A = 123
		te.b = "NOTUSED"
		return &te
	}
	var res = make([]*Test, b.N)
	p := NewUnsafePool(f)
	for j := 0; j < b.N; j++ {
		for i := 0; i < b.N; i++ {
			res[i] = p.Get()
		}
		for i := 0; i < b.N; i++ {
			res[i] = p.Get()
		}
	}
	x = res[rand.Int()%len(res)]
}

func Benchmark_PoolGetReturn(b *testing.B) {
	f := func() *Test {
		te := Test{}
		te.A = 123
		te.b = "NOTUSED"
		return &te
	}
	var res = make([]*Test, b.N)
	p := NewUnsafePool(f)
	for _i := 0; _i < b.N; _i++ {
		for i := 0; i < b.N; i++ {
			res[i] = p.Get()
		}
		for i := 0; i < b.N; i++ {
			p.Return(res[i])
		}
		for i := 0; i < b.N; i++ {
			res[i] = p.Get()
		}
	}
	x = res[rand.Int()%len(res)]
}

type Test struct {
	b  string
	A  int8
	b2 string
	b3 string
	b4 string
	C  float64
	B  []byte
}

// GOGC=off test with this
func Test_Pool(t *testing.T) {
	f := func() *Test {
		te := Test{}
		te.A = 123
		te.b = "NOTUSED"
		return &te
	}
	rand.Seed(time.Now().UnixNano())
	p := NewUnsafePool(f)
	for i := 0; i < 1000000000; i++ {
		res := p.Get()
		if res == nil {
			//fmt.Println("NILLLLLLLLLLL")
		}
		//fmt.Println("RES=",res)
		resV := res
		//fmt.Println(resV)
		resV.A = 123
		resV.b = "USED"
		v := resV
		tx := v
		_ = tx
		//if tx.A > rand.Intn(500000-123+1)+123 {
		// //fmt.Println("YES")
		//}
		//if len(resV.B) < 51 {
		// for z := 0; z < 2048; z++ {
		//  resV.B = append(resV.B, byte(z))
		// }
		//}
		//fmt.Println(resV)
		//runtime.KeepAlive(&resV)
		if i == 1 {
			//p.Return(resV)
		}
	}
	//res = p.Get()
	//resV = (*res).(Test)
	//fmt.Println(resV)
	//i++
}
