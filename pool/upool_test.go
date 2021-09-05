package pool

import (
	"fmt"
	"testing"
	"unsafe"
)

type PoolI interface {
	Get() interface{}
	Return(interface{})
}

type UnsafePool struct {
	New     func() interface{}
	sz, ptr uintptr
}

func NewUnsafePool(new func() interface{}) PoolI {
	p := &UnsafePool{
		New: new,
		sz:  0,
		ptr: 0,
	}
	return p
}

func (p *UnsafePool) Get() interface{} {
	var st interface{}
	if p.sz == 0 {
		st = p.New()
		p.sz = unsafe.Sizeof(st)
		p.ptr = uintptr(unsafe.Pointer(&st))
	} else {

	}
	return *(*interface{})(unsafe.Pointer(p.ptr))
}

func (p *UnsafePool) Return(interface{}) {

}

func Benchmark_PoolGet(b *testing.B) {
	for i := 0; i < b.N; i++ {
	}
}

func Test_Pool(t *testing.T) {
	f := func() interface{} {
		return 123
	}
	p := NewUnsafePool(f)
	res := p.Get().(int)
	fmt.Println(res)
}
