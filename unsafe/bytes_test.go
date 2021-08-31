package unsafe

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

//Benchmark_B2S 1000000000          0.3140 ns/op - b2s
//Benchmark string([]byte{...}) 17811870 57.49 ns/op
func Benchmark_B2S(b *testing.B) {
	by := []byte("ABCDEffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffsdaewwqre324567890-=-0987654321`23435678907870--9=-")
	for i := 0; i < b.N; i++ {
		_ = b2s(by)
	}
}

//Benchmark_S2B 1000000000          0.5691 ns/op
//Benchmark []byte(...) 16965518 61.76 ns/op
func Benchmark_S2B(b *testing.B) {
	s := "ABCDEffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffsdaewwqre324567890-=-0987654321`23435678907870--9=-"
	for i := 0; i < b.N; i++ {
		_ = s2b(s)
	}
}

func Test_B2S(t *testing.T) {
	s := "ABCDEffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffsdaewwqre324567890-=-0987654321`23435678907870--9=-"
	by := []byte(s)
	if s != b2s(by) {
		panic("s != b2s(s)")
	}
}

func Test_S2B(t *testing.T) {
	s := "ABCDEffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffsdaewwqre324567890-=-0987654321`23435678907870--9=-"
	if string(s2b(s)) != s {
		panic("s2b(s) != []byte(s)")
	}
}