package set

import "testing"

// declaration of memory optimized set: type Set map[T]struct{}

// IntSet ex: integer set with 0 bytes for value
type set map[int]struct{}

func Test_IntSet(t *testing.T) {
	s := set{}
	for i := 0; i < 50000; i++ {
		s[i] = struct{}{}
	}
	if _, ok := s[11]; !ok {
		t.Error("key 11 is not present")
	}

	if _, ok := s[50001]; ok {
		t.Error("key 50001 is present")
	}
}
