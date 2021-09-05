package pool

import (
	"fmt"
	"math/rand"
	"sync/atomic"
	"testing"
	"time"
	"unsafe"
)

type PoolI interface {
	Get() interface{}
	Return(*interface{})
}

type UnsafePool struct {
	New      func() interface{}
	sz       uintptr
	ptrs     []uintptr
	freeIdxs []int
}

func NewUnsafePool(new func() interface{}) *UnsafePool {
	p := &UnsafePool{
		New:      new,
		sz:       0,
		ptrs:     make([]uintptr, 0),
		freeIdxs: make([]int, 0),
	}
	return p
}

func (p *UnsafePool) uintptr2EmptyI(ptr uintptr) interface{} {
	return (*interface{})(unsafe.Pointer(ptr))
}

func (p *UnsafePool) emptyI2uintptr(ei interface{}) uintptr {
	return uintptr(unsafe.Pointer(ei.(*unsafe.ArbitraryType)))
}

func (p *UnsafePool) Get() interface{} {
	var st interface{}
	if len(p.ptrs) == 0 {
		st = p.New()
		p.sz = unsafe.Sizeof(st)
		p.ptrs = append(p.ptrs, p.emptyI2uintptr(st))
	} else {
		if len(p.freeIdxs) != 0 {
			i := p.freeIdxs[len(p.freeIdxs)-1]
			p.freeIdxs = p.freeIdxs[:len(p.freeIdxs)-1]
			st = p.uintptr2EmptyI(p.ptrs[i])
		} else {
			nextPtr := p.ptrs[len(p.ptrs)-1] + p.sz
			st = p.New()
			atomic.StoreUintptr(&nextPtr, p.emptyI2uintptr(st))
			p.ptrs = append(p.ptrs, nextPtr)
			st = p.uintptr2EmptyI(nextPtr)
		}
	}
	return st
}

func (p *UnsafePool) Return(st interface{}) {
	ptr := p.emptyI2uintptr(st)
	for i := range p.ptrs {
		if p.ptrs[i] == ptr {
			p.freeIdxs = append(p.freeIdxs, i)
		}
	}
}

var x int

func Benchmark_PoolGet(b *testing.B) {
	f := func() interface{} {
		return Test{}
	}
	p := NewUnsafePool(f)
	for j := 0; j < b.N; j++ {
		res := p.Get()
		resV := (res).(*Test)
		//fmt.Println(resV)
		resV.A = 123
		resV.b = "dsgfdfgdfsag"
		//fmt.Println(resV)
		p.Return(res)
	}
}

type Test struct {
	b string
	A int
	b2 string
	b3 string
	b4 string
	C float64
	B []byte
}

func Test_Pool(t *testing.T) {
	f := func() interface{} {
		te:=Test{}
		te.A = rand.Intn(500000 - 123 + 1) + 123
		te.b = "dsgfdfgdfsag"
		return &te
	}
	rand.Seed(time.Now().UnixNano())
	p := NewUnsafePool(f)
	for i:=0;i<1000000000;i++ {
		res := p.Get()
		resV := (res).(*Test)
		fmt.Println(resV)
		resV.A = rand.Intn(500000 - 123 + 1) + 123
		resV.b = "dsgfdfgdfsag"
		v := resV
		tx := v
		if tx.A > rand.Intn(500000 - 123 + 1) + 123 {
			fmt.Println("YES")
		}
		fmt.Println(resV)
		//runtime.KeepAlive(&resV)
		p.Return(res)
	}
	//res = p.Get()
	//resV = (*res).(Test)
	//fmt.Println(resV)
	//i++
}
