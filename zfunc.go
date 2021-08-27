package zfunc

import "fmt"

func main() {
	fmt.Println(zf("abacab", 0))
	// {0,0,1,0,2,1}
	fmt.Println(zf("aaaaa", 0))
	// {0,4,3,2,1}
}

func zf(s string, zv int) []int {
	li := 0
	ri := 0
	z := make([]int, len(s))
	z[0] = zv
	for i := 1; i < len(s); i++ {
		if i <= ri {
			x := ri - i + 1
			if z[i-li] < x {
				z[i] = z[i-li]
			} else {
				z[i] = x
			}
		}
		for i+z[i] < len(s) && s[z[i]] == s[i+z[i]] {
			z[i]++
		}
		if i+z[i]-1 > ri {
			li = 1
			ri = i + z[i] - 1
		}
	}
	return z
}
