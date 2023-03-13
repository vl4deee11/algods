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

func con(s string) uint32 {
	d := uint32(0)
	ll := len(s) - 1
	for i := 0; i < len(s); i++ {
		d += uint32(s[i] - byte(48))
		if i < ll {
			d *= 10
		}
	}

	return d
}
