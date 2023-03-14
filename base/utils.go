package base

import (
	"bufio"
)

type ii int
type aii []int
type mii [][]int
type i3 int32
type ai3 []int32
type mi3 [][]int32
type i6 int64
type ai6 []int64
type mi6 [][]int64
type i1 int16
type ai1 []int16
type mi1 [][]int16
type i8 int8
type ai8 []int8
type mi8 [][]int8
type ui uint
type aui []uint
type mui [][]uint
type u3 uint32
type au3 []uint32
type mu3 [][]uint32
type u6 uint64
type au6 []uint64
type mu6 [][]uint64
type u1 uint16
type au1 []uint16
type mu1 [][]uint16
type u8 uint8
type au8 []uint8
type mu8 [][]uint8
type f3 float32
type af3 []float64
type mf3 [][]float64
type f6 float64
type af6 []float64
type mf6 [][]float64
type c6 complex64
type c1 complex128
type by byte
type aby []byte
type mby [][]byte
type bo bool
type abo []bool
type mbo [][]bool
type s string
type as []string
type ms [][]string
type r rune
type ar []rune
type mr [][]rune

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
