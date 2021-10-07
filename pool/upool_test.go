package pool

import (
	"testing"
	"unsafe"
)

type Test struct {
	b  string
	bx string
	A  int8
	b2 string
	b3 string
	b4 string
	C  float64
	B  []byte
}

const maxInChunks = 100
const size = int(unsafe.Sizeof(Test{}))

type UPool struct {
	New        func() Test
	currChunk  int
	currOffset int
	freeptrs   []uintptr
	memChunks  [][maxInChunks * size]byte
}

func NewUPool(new func() Test, pSize int) *UPool {
	return &UPool{
		New:       new,
		memChunks: make([][maxInChunks * size]byte, 0, pSize),
		freeptrs:  make([]uintptr, 0, pSize),
	}
}

// noescape: unused now
func (p *UPool) noescape(up unsafe.Pointer) unsafe.Pointer {
	return unsafe.Pointer(uintptr(up) ^ 0)
}

// uintptr2t: uintptr to type convert
func (p *UPool) uintptr2t(ptr uintptr) *Test {
	return (*Test)(unsafe.Pointer(ptr))
}

// t2uintptr: type to uintptr convert
func (p *UPool) t2uintptr(v *Test) uintptr {
	return uintptr(unsafe.Pointer(v))
}

// malloc: allocate new memory chunk
func (p *UPool) malloc() {
	p.memChunks = append(p.memChunks, [maxInChunks * size]byte{})
}

// Get struct from pool
func (p *UPool) Get() *Test {
	if len(p.freeptrs) != 0 {
		ptr := p.freeptrs[len(p.freeptrs)-1]
		p.freeptrs = p.freeptrs[:len(p.freeptrs)-1]
		return p.uintptr2t(ptr)
	}

	st := p.New()
	if len(p.memChunks) == 0 {
		p.malloc()
	}

	if p.currOffset == len(p.memChunks[p.currChunk]) {
		p.malloc()
		p.currOffset = 0
		p.currChunk++
	}

	ptr := unsafe.Pointer(&st)

	bs := *(*[size]byte)(ptr)
	for i := range bs {
		p.memChunks[p.currChunk][p.currOffset+i] = bs[i]
	}
	p.currOffset += size
	return (*Test)(unsafe.Pointer(&p.memChunks[p.currChunk][p.currOffset-size]))
}

// Return struct to pool
func (p *UPool) Return(st *Test) {
	p.freeptrs = append(p.freeptrs, p.t2uintptr(st))
}

var x *Test

// go test -bench=. -gcflags '-l -N' -benchmem -cpu=1
// goos: linux
// goarch: amd64
// Benchmark_PoolGetOnly           10000000               759.1 ns/op           627 B/op          1 allocs/op
// Benchmark_PoolGetReturn         10000000                12.23 ns/op            0 B/op          0 allocs/op
func Benchmark_PoolGetOnly(b *testing.B) {
	f := func() Test {
		te := Test{}
		te.A = 123
		te.b = "NOTUSED"
		te.B = []byte{1, 2, 3, 4}
		return te
	}

	var v *Test
	p := NewUPool(f, 1000)
	for j := 0; j < b.N; j++ {
		v = p.Get()
	}
	x = v
}

func Benchmark_PoolGetReturn(b *testing.B) {
	f := func() Test {
		te := Test{}
		te.A = 123
		te.b = "NOTUSED"
		te.B = []byte{1, 2, 3, 4}
		return te
	}

	var v *Test
	p := NewUPool(f, 1000)
	for _i := 0; _i < b.N; _i++ {
		v = p.Get()
		p.Return(v)
	}
	x = v
}

func Test_Pool(t *testing.T) {
	f := func() Test {
		te := Test{}
		te.A = 123
		te.b = "NOTUSED"
		te.B = []byte{1, 2, 3, 4}
		return te
	}

	p := NewUPool(f, 10)
	var objects []*Test
	for i := 0; i < 10000; i++ {
		res := p.Get()
		resV := res
		objects = append(objects, resV)
		if i%2 == 0 {
			p.Return(objects[i])
		}
	}
}
