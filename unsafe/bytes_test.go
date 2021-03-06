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

// uint64 to bit vector
func uInt642bitVec(b uint64) []uint8 {
	v := make([]uint8, 64)
	i := len(v) - 1
	for b != 0 {
		v[i] = uint8(b & 1)
		b >>= 1
		i--
	}
	return v
}

// bit vector to uint64
func bitVec2UInt64(v []uint8) uint64 {
	e := len(v) - 1
	if e > 63 {
		return 0
	}

	var b uint64 = 0
	for i := range v {
		if v[i] == 1 {
			b |= 1
		}
		if i != e {
			b <<= 1
		}
	}
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

func Test_IntToVecAndRollback(t *testing.T) {
	var i uint64 = 0
	for i = 0; i < 1<<24; i++ {
		r := bitVec2UInt64(uInt642bitVec(i))
		if r != i {
			t.Errorf("test: %d != %d", r, i)
		}
	}

	r := bitVec2UInt64(uInt642bitVec(1 << 63))
	if r != uint64(1<<63) {
		t.Errorf("test: %d != %d", r, uint64(1<<63))
	}
}
