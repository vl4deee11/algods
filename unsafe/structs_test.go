package ubytes

import (
	"fmt"
	"testing"
	"ubytes/pkg"
	"unsafe"
)

// https://hackernoon.com/golang-unsafe-type-conversions-and-memory-access-odz3yrl

func TestOffset(t *testing.T) {
	x := pkg.NewABC(3, 12, 5)
	pp := uintptr(unsafe.Pointer(x)) + 2
	v := (*int8)(unsafe.Pointer(pp))
	*v = 7
	fmt.Println(x)
}
