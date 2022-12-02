package string

import (
	"reflect"
	"testing"
)

// zf - z function algo with O(N)
// link to docs - http://e-maxx.ru/algo/z_function
func zf(s string, zv int) []int {
	if len(s) == 0 {
		return []int{}
	}
	li := 0
	ri := 0
	z := make([]int, len(s))
	z[0] = zv
	for i := 1; i < len(s); i++ {
		if i <= ri {
			x := ri - i
			if z[i-li] < x {
				z[i] = z[i-li]
			} else {
				z[i] = x
			}
		}
		for i+z[i] < len(s) && s[z[i]] == s[i+z[i]] {
			z[i]++
		}

		if i+z[i] > ri {
			li = i
			ri = i + z[i]
		}
	}
	return z
}

func Test_ZF(t *testing.T) {
	z1 := zf("abacab", 0)
	z2 := zf("cats$longcatss", 0)
	z3 := zf("", 0)
	z4 := zf("a", 0)
	z5 := zf("aaaaa", 5)
	if !reflect.DeepEqual(z1, []int{0, 0, 1, 0, 2, 0}) {
		t.Error("z1")
	}
	if !reflect.DeepEqual(z2, []int{0, 0, 0, 0, 0, 0, 0, 0, 0, 4, 0, 0, 0, 0}) {
		t.Error("z2")
	}
	if !reflect.DeepEqual(z3, []int{}) {
		t.Error("z3")
	}
	if !reflect.DeepEqual(z4, []int{0}) {
		t.Error("z4")
	}
	if !reflect.DeepEqual(z5, []int{5, 4, 3, 2, 1}) {
		t.Error("z5")
	}
}

var resS []int

// Benchmark_ZF   	 1974745	       598.3 ns/op - use raw string
// Benchmark_ZF   	 1307140	       958.8 ns/op - with string to rune
func Benchmark_ZF(b *testing.B) {
	s := "catsssssssssssssssssssssssssssssssss$longcatsssssssssssssssssssssssssssssssssssssssssssssssssssssssssssssssssscats"
	var r []int
	for i := 0; i < b.N; i++ {
		r = zf(s, 0)
	}
	resS = r
}
