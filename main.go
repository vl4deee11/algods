package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
)

var stdinR = bufio.NewReaderSize(os.Stdin, 1<<30)
var stdoutW = bufio.NewWriterSize(os.Stdout, 1<<30)

// Если нужно много писать
//var stdoutW = bufio.NewWriterSize(os.Stdout, 1<<30)

func printf(f string, a ...interface{})             { fmt.Fprintf(stdoutW, f, a...) }
func print(a ...interface{})                        { fmt.Fprint(stdoutW, a...) }
func println(a ...interface{})                      { fmt.Fprintln(stdoutW, a...) }
func scanf(f string, a ...interface{}) (int, error) { return fmt.Fscanf(stdinR, f, a...) }
func scan(a ...interface{})                         { fmt.Fscan(stdinR, a...) }
func scanln(a ...interface{})                       { fmt.Fscanln(stdinR, a...) }
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

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

func strToI(s string) int {
	x, _ := strconv.Atoi(s)
	return x
}

func chazToI026(x byte) int {
	return int(x - 97)
}

func chAZToI026(x byte) int {
	return int(x - 65)
}

func i026ToChaz(x int) byte {
	return byte(x + 97)
}

func i026ToChAZ(x int) byte {
	return byte(x + 65)
}

func isLet(b byte) bool {
	return (97 <= b && b <= 122) || (65 <= b && b <= 90)
}

func isN(b byte) bool {
	return 48 <= b && b <= 57
}

func readFullLine(r *bufio.Reader) string {
	l, _ := r.ReadString('\n')
	return l[:len(l)-1]
}

func dist(p1 [2]int, p2 [2]int) float64 {
	return math.Sqrt(math.Pow(float64(p2[0]-p1[0]), 2) + math.Pow(float64(p2[1]-p1[1]), 2))
}

func mod(a, b int) int {
	return (a%b + b) % b
}

func chess(b string) [2]int {
	return [2]int{7 - ((int(b[0]) - 48) - 1), 7 - (int(b[1]) - 97)}
}

func float64Eq(f1, f2 float64) bool {
	return math.Abs(f1-f2) < 1e-6
}

func float64GtOrEq(f1, f2 float64) bool {
	return float64Eq(f1, f2) || f1 > f2
}

func float64LtOrEq(f1, f2 float64) bool {
	return float64Eq(f1, f2) || f1 < f2
}

func ch(a, b, c [2]int64) bool {
	return int64(b[0]-a[0])*int64(c[1]-a[1])-int64(b[1]-a[1])*int64(c[0]-a[0]) >= 0
}

type Heap []int

func (h Heap) Len() int            { return len(h) }
func (h Heap) Less(i, j int) bool  { return h[i] < h[j] }
func (h Heap) Swap(i, j int)       { h[i], h[j] = h[j], h[i] }
func (h *Heap) Push(x interface{}) { *h = append(*h, x.(int)) }
func (h *Heap) Pop() (v interface{}) {
	*h, v = (*h)[:len(*h)-1], (*h)[len(*h)-1]
	return
}
func gcd(a, b int64) int64 {
	if b == 0 {
		return a
	}

	return gcd(b, a%b)
}
func sumAp(a, b int64) int64 {
	n := (b - a) + 1
	return (n * (a + b)) / 2
}

type ufmst struct {
	p, r []int
}

func (uf *ufmst) find(e int) int {
	if uf.p[e] == e {
		return e
	}
	uf.p[e] = uf.find(uf.p[e])
	return uf.p[e]
}

func (uf *ufmst) union(e1, e2 int) {
	r1 := uf.find(e1)
	r2 := uf.find(e2)
	if r1 == r2 {
		return
	}

	switch {
	case uf.r[r1] < uf.r[r2]:
		uf.p[r1] = r2
	case uf.r[r1] > uf.r[r2]:
		uf.p[r2] = r1
	default:
		uf.p[r2] = r1
		uf.r[r1]++
	}
	return
}

// cat i.txt | go run main.go > o.txt
func main() {
	defer stdoutW.Flush()

}

func isPerfectSquare(x int64) bool {
	s := int64(math.Sqrt(float64(x)))
	return s*s == x
}
