package math

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// binpow - за O(log N)
// Бинарное (двоичное) возведение в степень — это приём, позволяющий возводить любое число в n-ую степень за O(log N)
// умножений (вместо n умножений при обычном подходе).
// Заметим что для четной степени a^n = a^(n/2)*a^(n/2), а нечетную степень к четной можно привести через вычитание единицы
func binpow(a, n int) int {
	if n == 0 {
		return 1
	}
	if n&1 != 0 {
		return binpow(a, n-1) * n
	}
	x := binpow(a, n/2)
	return x * x
}

func TestBinpow(t *testing.T) {
	assert.Equal(t, binpow(3, 3), 27)
	assert.Equal(t, binpow(2, 3), 8)
	assert.Equal(t, binpow(2, 8), 1<<8)
}
