package zfunc

import (
	"reflect"
	"testing"
)

// zf - за время O(N)
//       a b a c a b a
// z =  [7 0 1 0 3 0 1]
//       a a a a a
// z =  [0 4 3 2 1]
//       a b c d
// z =  [0 0 0 0]
//       a c f a c f a c
// z =  [0 0 0 5 0 0 2 0]
// pi = 0; pi = 1; pi = 2; pi = 3; pi = 4;
// ci = 3; ci = 4; ci = 5; ci = 6; ci = 7; = 5
// z[i] = max совпадение префикса строки и строки начиная с i символа строки
// s =  c a t s $ l o n g c a t s s
// z = [0 0 0 0 0 0 0 0 0 4 0 0 0 0]
//
// s = car
// t = my favorite carting

func zf(s string, zv int) []int {
	if len(s) == 0 {
		return []int{}
	}
	li := 0
	ri := 0
	z := make([]int, len(s))
	z[0] = zv
	for i := 1; i < len(s); i++ {
		// Если i попадает в кещ до ri
		if i <= ri {
			// х - то сколько мы можем взять из кэша
			x := ri - i
			if z[i-li] < x {
				// Если закешированно меньше чем осталось в кеше
				z[i] = z[i-li]
			} else {
				// Если закешированно больше или равно чем осталось в кеше
				z[i] = x
			}
		}

		// Досчитывание в тупую
		for i+z[i] < len(s) && s[z[i]] == s[i+z[i]] {
			z[i]++
		}

		// Обновление границ, для того что бы следующие i индексы попали в кеш
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
