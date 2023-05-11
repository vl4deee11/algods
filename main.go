package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"sort"
)

var stdinR = bufio.NewReader(os.Stdin)
var stdoutW = bufio.NewWriter(os.Stdout)

func printf(f string, a ...interface{}) { fmt.Fprintf(stdoutW, f, a...) }
func print(a ...interface{})            { fmt.Fprint(stdoutW, a...) }
func println(a ...interface{})          { fmt.Fprintln(stdoutW, a...) }
func scanf(f string, a ...interface{})  { fmt.Fscanf(stdinR, f, a...) }
func scan(a ...interface{})             { fmt.Fscan(stdinR, a...) }
func scanln(a ...interface{})           { fmt.Fscanln(stdinR, a...) }
func strToI(s string) int {
	m := false
	if s[0] == '-' {
		m = true
		s = s[1:]
	}
	d := int(0)
	ll := len(s) - 1
	for i := 0; i < len(s); i++ {
		d += int(s[i] - byte(48))
		if i < ll {
			d *= 10
		}
	}
	if m {
		return -d
	}
	return d
}
func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}

func min(x, y int) int {
	if x < y {
		return x
	}
	return y
}

func con(s string) int {
	d := 0
	ll := len(s) - 1
	for i := 0; i < len(s); i++ {
		d += int(s[i] - byte(48))
		if i < ll {
			d *= 10
		}
	}

	return d
}
func isN(b byte) bool {
	return 48 <= b && b <= 57
}
func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}
func isLet(b byte) bool {
	return (97 <= b && b <= 122) || (65 <= b && b <= 90)
}

var tr = make(map[byte]interface{})

func main() {
	defer stdoutW.Flush()

	var t int
	scanf("%d\n", &t)
	for ti := 0; ti < t; ti++ {
		var n int
		scanf("%d\n", &n)
		var ns = make([]string, n)
		for i := range ns {
			scanf("%s\n", &ns[i])
		}
		sort.Slice(ns, func(i, j int) bool {
			return len(ns[i]) < len(ns[j])
		})
		tr = make(map[byte]interface{})
		f := false
	ml:
		for i := range ns {
			cp := tr
			for j := range ns[i] {
				if _, ok := cp[ns[i][j]]; !ok {
					cp[ns[i][j]] = make(map[byte]interface{})
				}
				if _, ok := cp['$']; ok {
					f = true
					break ml
				}
				cp = cp[ns[i][j]].(map[byte]interface{})
			}
			cp['$'] = true
		}
		if f {
			println("NO")
		} else {
			println("YES")
		}
	}

}

func EuclideanDist(p1 [2]int, p2 [2]int) int {
	return int(math.Pow(float64(p2[0]-p1[0]), 2) + math.Pow(float64(p2[1]-p1[1]), 2))
}

func float64Eq(f1, f2 float64) bool {
	return math.Abs(f1-f2) < 1e-4
}

func float64GtOrEq(f1, f2 float64) bool {
	return float64Eq(f1, f2) || f1 > f2
}

func float64LtOrEq(f1, f2 float64) bool {
	return float64Eq(f1, f2) || f1 < f2
}
