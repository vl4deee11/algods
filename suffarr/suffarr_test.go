package suffarr

import (
	"reflect"
	"sort"
	"testing"
)

type spair struct {
	p iipair
	s int
}

type iipair struct {
	i int
	j int
}

type bipair struct {
	b byte
	i int
}

func clear(a []int) {
	for i := range a {
		a[i] = 0
	}
}

func radixSort(a []spair) {
	n := len(a)
	cnt := make([]int, n)
	pos := make([]int, n)
	r := make([]spair, n)
	{
		// second element in pair
		for i := range a {
			cnt[a[i].p.j]++
		}

		for i := 1; i < n; i++ {
			pos[i] = pos[i-1] + cnt[i-1]
		}

		for _, sp := range a {
			x := sp.p.j
			r[pos[x]] = sp
			pos[x]++
		}
		copy(a, r)
	}

	{
		// first element in pair
		clear(cnt)
		for i := range a {
			cnt[a[i].p.i]++
		}

		for i := range r {
			r[i] = spair{}
		}
		clear(pos)
		for i := 1; i < n; i++ {
			pos[i] = pos[i-1] + cnt[i-1]
		}

		for _, sp := range a {
			x := sp.p.i
			r[pos[x]] = sp
			pos[x]++
		}
		copy(a, r)
	}
}

// suffixArr - suffix array algo with O(NLogN)
// link to docs - http://e-maxx.ru/algo/suffix_array
func suffixArr(s string) []int {
	n := len(s)
	p := make([]int, n) // res
	c := make([]int, n) // class equivalent

	{
		// k = 0
		a := make([]*bipair, n)
		for i := 0; i < n; i++ {
			a[i] = &bipair{s[i], i}
		}
		sort.Slice(a, func(i, j int) bool {
			if a[i].b != a[j].b {
				return a[i].b < a[j].b
			}
			return a[i].i < a[j].i
		})
		for i := 0; i < n; i++ {
			p[i] = a[i].i
		}
		c[p[0]] = 0
		for i := 1; i < n; i++ {
			if a[i-1].b == a[i].b {
				c[p[i]] = c[p[i-1]]
			} else {
				c[p[i]] = c[p[i-1]] + 1
			}
		}
	}

	k := 0
	// 1 << k == 2**k
	for (1 << k) < n {
		// k -> k + 1
		a := make([]spair, n)
		for i := 0; i < n; i++ {
			a[i] = spair{
				p: iipair{
					i: c[i],
					j: c[(i+(1<<k))%n],
				},
				s: i,
			}
		}
		radixSort(a)

		for i := 0; i < n; i++ {
			p[i] = a[i].s
		}

		c[p[0]] = 0

		for i := 1; i < n; i++ {
			if a[i-1].p.i == a[i].p.i && a[i-1].p.j == a[i].p.j {
				c[p[i]] = c[p[i-1]]
			} else {
				c[p[i]] = c[p[i-1]] + 1
			}
		}
		k++
	}

	// show arr with strs
	//for i := 0; i < n; i++ {
	//	fmt.Println(p[i], s[p[i]:n])
	//}
	return p
}

func Test_SuffixArr(t *testing.T) {
	// $ - spec divisor
	s := "ababba$"
	sa1 := suffixArr(s)
	if !reflect.DeepEqual(sa1, []int{6, 5, 0, 2, 4, 1, 3}) {
		t.Error("sa1")
	}

	s2 := "aaaaa$"
	sa2 := suffixArr(s2)
	if !reflect.DeepEqual(sa2, []int{5, 4, 3, 2, 1, 0}) {
		t.Error("sa2")
	}

	s3 := "ppppplppp$"
	sa3 := suffixArr(s3)
	if !reflect.DeepEqual(sa3, []int{9, 5, 8, 4, 7, 3, 6, 2, 1, 0}) {
		t.Error("sa3")
	}

}
