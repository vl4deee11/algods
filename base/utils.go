package base

import (
	"bufio"
	"math"
)

type pii = [2]int
type mpii = [][2]int
type ii = int
type aii = []int
type mii = [][]int
type i3 = int32
type ai3 = []int32
type mi3 = [][]int32
type i6 = int64
type ai6 = []int64
type mi6 = [][]int64
type i1 = int16
type ai1 = []int16
type mi1 = [][]int16
type i8 = int8
type ai8 = []int8
type mi8 = [][]int8
type ui = uint
type aui = []uint
type mui = [][]uint
type u3 = uint32
type au3 = []uint32
type mu3 = [][]uint32
type u6 = uint64
type au6 = []uint64
type mu6 = [][]uint64
type u1 = uint16
type au1 = []uint16
type mu1 = [][]uint16
type u8 = uint8
type au8 = []uint8
type mu8 = [][]uint8
type f3 = float32
type af3 = []float64
type mf3 = [][]float64
type f6 = float64
type af6 = []float64
type mf6 = [][]float64
type c6 = complex64
type c1 = complex128
type by = byte
type aby = []byte
type mby = [][]byte
type bo = bool
type abo = []bool
type mbo = [][]bool
type s = string
type as = []string
type ms = [][]string
type r = rune
type ar = []rune
type mr = [][]rune

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
	d := int(0)
	ll := len(s) - 1
	for i := 0; i < len(s); i++ {
		d += int(s[i] - byte(48))
		if i < ll {
			d *= 10
		}
	}

	return d
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

func bytes2Int64(v []byte) int64 {
	if len(v) == 0 {
		return 0
	}
	if (len(v) == 8 && v[0] > 127) || len(v) > 8 {
		return 0
	}
	var b int64
	for i := 0; i < len(v); i++ {
		b |= int64(v[i])
		if b == 0 {
			b = 1
		}
		if len(v)-1 != i {
			b = b << 8
		}
	}
	if v[0] == 0 {
		b = b & ((1 << ((len(v) - 1) * 8)) - 1)
	}
	return b
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
