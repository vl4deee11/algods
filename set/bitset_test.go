package set

import (
	"testing"
)

type BMSet struct {
	masks []uint64
	zero  bool
}

type SetI interface {
	Set(i int)
	Get(i int) bool
	Delete(i int)
	Intersect(oth *BMSet) SetI
	Union(oth *BMSet) SetI
}

func New(size uint64) SetI {
	c := size / 64
	if size%64 != 0 {
		c++
	}
	return &BMSet{
		masks: make([]uint64, c),
	}
}

func (s *BMSet) Set(i int) {
	if i == 0 {
		s.zero = true
		return
	}
	bn, k := s.getSettings(i)
	s.masks[bn] = s.masks[bn] | k
}

func (s *BMSet) Get(i int) bool {
	if i == 0 {
		return s.zero
	}
	bn, k := s.getSettings(i)
	return s.masks[bn]&k != 0
}

func (s *BMSet) Delete(i int) {
	if i == 0 {
		s.zero = false
		return
	}
	bn, k := s.getSettings(i)
	s.masks[bn] = s.masks[bn] & (^k)
}

func (s *BMSet) getSettings(i int) (int, uint64) {
	bn := i / 64
	if i%64 != 0 {
		bn++
	}
	if bn > 0 {
		bn--
	}

	return bn, uint64(1 << (i % 64))
}

func (s *BMSet) Intersect(oth *BMSet) SetI {
	ll := len(oth.masks)
	if len(s.masks) > ll {
		ll = len(s.masks)
	}
	masks := make([]uint64, ll)
	for i := 0; i < ll; i++ {
		if i < len(s.masks) && i < len(oth.masks) {
			masks[i] = s.masks[i] | oth.masks[i]
		} else if i < len(s.masks) {
			masks[i] = s.masks[i]
		} else if i < len(oth.masks) {
			masks[i] = oth.masks[i]
		}
	}
	if s.zero || oth.zero {
		return &BMSet{zero: true, masks: masks}
	}
	return &BMSet{zero: false, masks: masks}
}

func (s *BMSet) Union(oth *BMSet) SetI {
	ll := len(oth.masks)
	if len(s.masks) < ll {
		ll = len(s.masks)
	}
	masks := make([]uint64, ll)
	for i := 0; i < ll; i++ {
		masks[i] = s.masks[i] & oth.masks[i]
	}
	if s.zero && oth.zero {
		return &BMSet{zero: true, masks: masks}
	}
	return &BMSet{zero: false, masks: masks}
}






// BenchmarkBMSet        100000000               10.6 ns/op             0 B/op          0 allocs/op
func BenchmarkBMSet(b *testing.B) {
	sz := uint64(b.N)
	bms := New(sz)
	for i := 0; i < b.N; i++ {
		bms.Set(i)
		v := bms.Get(i)
		if v {

		}
		bms.Delete(i)
	}
}

func TestBMSet1(t *testing.T) {
	sz := 131
	bms := New(uint64(sz))
	for i := 0; i <= sz; i++ {
		bms.Set(i)
		if !bms.Get(i) {
			t.Errorf("test !bms.Get(i) fail on = %d", i)
			return
		}
	}

	// Delete x & 1 == 0
	for i := 0; i <= sz; i++ {
		if !bms.Get(i) {
			t.Errorf("test !bms.Get(i) fail on = %d", i)
			return
		}
		if i&1 == 0 {
			bms.Delete(i)
			if bms.Get(i) {
				t.Errorf("test bms.Get(i) fail on = %d", i)
				return
			}
		}
	}

	// Delete x & 1 != 0
	for i := 0; i <= sz; i++ {
		if i&1 != 0 {
			if !bms.Get(i) {
				t.Errorf("test !bms.Get(i) fail on = %d", i)
				return
			}
			bms.Delete(i)
			if bms.Get(i) {
				t.Errorf("test bms.Get(i) fail on = %d", i)
				return
			}
		} else {
			if bms.Get(i) {
				t.Errorf("test bms.Get(i) fail on = %d", i)
				return
			}
		}
	}

	for i := 0; i <= sz; i++ {
		if bms.Get(i) {
			t.Errorf("test bms.Get(i) fail on = %d", i)
			return
		}
	}
}

func TestBMSet2(t *testing.T) {
	sz := 128
	bms := New(uint64(sz))
	for i := 0; i <= sz; i++ {
		bms.Set(i)
		if !bms.Get(i) {
			t.Errorf("test !bms.Get(i) fail on = %d", i)
			return
		}
	}

	// Delete x & 1 == 0
	for i := 0; i <= sz; i++ {
		if !bms.Get(i) {
			t.Errorf("test !bms.Get(i) fail on = %d", i)
			return
		}
		if i&1 == 0 {
			bms.Delete(i)
			if bms.Get(i) {
				t.Errorf("test bms.Get(i) fail on = %d", i)
				return
			}
		}
	}

	// Delete x & 1 != 0
	for i := 0; i <= sz; i++ {
		if i&1 != 0 {
			if !bms.Get(i) {
				t.Errorf("test !bms.Get(i) fail on = %d", i)
				return
			}
			bms.Delete(i)
			if bms.Get(i) {
				t.Errorf("test bms.Get(i) fail on = %d", i)
				return
			}
		} else {
			if bms.Get(i) {
				t.Errorf("test bms.Get(i) fail on = %d", i)
				return
			}
		}
	}

	for i := 0; i <= sz; i++ {
		if bms.Get(i) {
			t.Errorf("test bms.Get(i) fail on = %d", i)
			return
		}
	}
}

func TestIntersect(t *testing.T) {
	bm1 := New(128)
	bm1.Set(3)
	bm1.Set(65)
	bm1.Set(120)
	bm1.Set(0)
	bm1.Set(4)

	bm2 := New(67)
	bm2.Set(3)
	bm2.Set(66)

	bm3 := bm1.Intersect(bm2.(*BMSet))
	vals := []int{3, 65, 0, 4, 66, 120}
	for i := 0; i < len(vals); i++ {
		if !bm3.Get(vals[i]) {
			t.Errorf("test bms.Get(i) fail on = %d", vals[i])
			return
		}
		bm3.Delete(vals[i])
	}

	for i := 0; i < 128; i++ {
		if bm3.Get(i) {
			t.Errorf("test bms.Get(i) fail on = %d", i)
			return
		}
	}
}

func TestUnion(t *testing.T) {
	bm1 := New(128)
	bm1.Set(3)
	bm1.Set(66)
	bm1.Set(120)
	bm1.Set(0)
	bm1.Set(4)

	bm2 := New(67)
	bm2.Set(3)
	bm2.Set(66)

	bm3 := bm1.Union(bm2.(*BMSet))
	vals := []int{3, 66}
	for i := 0; i < len(vals); i++ {
		if !bm3.Get(vals[i]) {
			t.Errorf("test bms.Get(i) fail on = %d", vals[i])
			return
		}
		bm3.Delete(vals[i])
	}

	for i := 0; i < 67; i++ {
		if bm3.Get(i) {
			t.Errorf("test bms.Get(i) fail on = %d", i)
			return
		}
	}
}