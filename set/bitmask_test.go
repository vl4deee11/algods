package set

import (
	"fmt"
	"testing"
)

func In(bm, e uint64) bool {
	return bm&e != 0
}

func Add(bm, e uint64) uint64 {
	return bm | e
}

func Del(bm, e uint64) uint64 {
	return bm & ^e
}

func PtrAdd(bm *uint64, e uint64) {
	*bm = *bm | e
}

func PtrDel(bm *uint64, e uint64) {
	*bm = *bm & ^e
}

func Test_IntBitMaskNoPtr(t *testing.T) {
	var bm uint64 = 0
	for i := 0; i < 64; i++ {
		bm = Add(bm, 1<<i)
	}

	ss := fmt.Sprintf("%b", bm)
	fmt.Println(ss)
	ks := []uint64{12, 63, 0, 11, 46, 37}
	for i := range ks {
		kk := uint64(1 << ks[i])
		fmt.Printf("Key=%b(%d)\n", kk, ks[i])
		fmt.Printf("Mask before deleting=%b\n", bm)
		if !In(bm, kk) {
			t.Errorf("key %d is not present", kk)
		}

		bm = Del(bm, kk)
		fmt.Printf("Mask after  deleting=%b\n", bm)

		if In(bm, kk) {
			t.Errorf("key %d is present", kk)
		}
	}
}

func Test_IntBitMaskPtr(t *testing.T) {
	var bm uint64 = 0
	for i := 0; i < 64; i++ {
		PtrAdd(&bm, 1<<i)
	}

	ss := fmt.Sprintf("%b", bm)
	fmt.Println(ss)
	ks := []uint64{12, 63, 0, 11, 46, 37}
	for i := range ks {
		kk := uint64(1 << ks[i])
		fmt.Printf("Key=%b(%d)\n", kk, ks[i])
		fmt.Printf("Mask before deleting=%b\n", bm)
		if !In(bm, kk) {
			t.Errorf("key %d is not present", kk)
		}

		PtrDel(&bm, kk)
		fmt.Printf("Mask after  deleting=%b\n", bm)

		if In(bm, kk) {
			t.Errorf("key %d is present", kk)
		}
	}
}

// go test -bench=. -gcflags '-l -N' -benchmem -cpu=1 -benchtime=1000000x
// Benchmark_IntBitMaskPtr         100000000               10.80 ns/op            0 B/op          0 allocs/op
// Benchmark_IntBitMaskNoPtr       100000000               12.25 ns/op            0 B/op          0 allocs/op
func Benchmark_IntBitMaskNoPtr(b *testing.B) {
	var bm uint64 = 0
	for i := 0; i < b.N; i++ {
		kk := uint64(1 << (i % 63))
		bm = Add(bm, kk)
	}

	for i := 0; i < b.N; i++ {
		kk := uint64(1 << (i % 63))
		In(bm, kk)
		bm = Del(bm, kk)
	}
}

func Benchmark_IntBitMaskPtr(b *testing.B) {
	var bm uint64 = 0
	for i := 0; i < b.N; i++ {
		kk := uint64(1 << (i % 63))
		PtrAdd(&bm, kk)
	}

	for i := 0; i < b.N; i++ {
		kk := uint64(1 << (i % 63))
		In(bm, kk)
		PtrDel(&bm, kk)
	}
}
