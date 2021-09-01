package ubytes

import (
	"reflect"
	"runtime"
	"testing"
	"unsafe"
)

// b2s converts byte slice to a string without memory allocation.
func b2s(b []byte) string {
	return *(*string)(unsafe.Pointer(&b))
}

// s2b converts string to a byte slice without memory allocation.
func s2b(s string) (b []byte) {
	bh := (*reflect.SliceHeader)(unsafe.Pointer(&b))
	sh := (*reflect.StringHeader)(unsafe.Pointer(&s))
	bh.Data = sh.Data
	bh.Cap = sh.Len
	bh.Len = sh.Len
	runtime.KeepAlive(&s)
	return b
}

var resS string
var resB []byte

//Benchmark_B2S 	1000000000	         0.8335 ns/op - b2s
//Benchmark string([]byte{...}) 	17811870	        57.49 ns/op
func Benchmark_B2S(b *testing.B) {
	by := []byte("ABCDEffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffsdaewwqre324567890-=-0987654321`23435678907870--9=-")
	r := ""
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		r = b2s(by)
	}
	resS = r
}

//Benchmark_S2B   	1000000000	         2.006 ns/op
//Benchmark []byte(...) 	16965518	        61.76 ns/op
func Benchmark_S2B(b *testing.B) {
	s := "ABCDEffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffsdaewwqre324567890-=-0987654321`23435678907870--9=-"
	var r []byte
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		r = s2b(s)
	}
	resB = r
}

func Test_B2S(t *testing.T) {
	s := "ABCDEffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffsdaewwqre324567890-=-0987654321`23435678907870--9=-"
	by := []byte(s)
	if s != b2s(by) {
		t.Error("s != b2s(s)")
	}
}

func Test_S2B(t *testing.T) {
	s := "ABCDEffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffsdaewwqre324567890-=-0987654321`23435678907870--9=-"
	if string(s2b(s)) != s {
		t.Error("s2b(s) != []byte(s)")
	}
}
