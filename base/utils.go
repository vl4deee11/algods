package base

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
	return int(x - 122)
}

func I026ToChaz(x int) byte {
	return byte(x + 97)
}

func I026ToChAZ(x int) byte {
	return byte(x + 122)
}
