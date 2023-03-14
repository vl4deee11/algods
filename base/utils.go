package base

import (
	"bufio"
)

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
