package math_x

import (
	"testing"

	"github.com/vl4deee11/bm_set"

	"github.com/stretchr/testify/assert"
)

// Решето Эратосфена — это алгоритм, позволяющий найти все простые числа в отрезке [1; n]
// Идея проста — запишем ряд чисел 1...n, и будем вычеркивать сначала все числа, делящиеся на 2,
// кроме самого числа 2, затем делящиеся на 3, кроме самого числа 3, затем на 5, затем на 7, 11,
// и все остальные простые до n.

// eratosthenesSieve - за  O (N*loglogN)
func eratosthenesSieve(n int) []bool {
	primes := make([]bool, n+1)
	for i := range primes {
		primes[i] = true
	}
	primes[1] = false
	primes[0] = false
	// Для того чтобы найти все простые до n, достаточно выполнить просеивание только простыми, не превосходящими корня из n.
	// Так как все числа просты в изначальном созданном массиве, то все числа до SQRT(N) будут обработаны с последующим выкл/вкл до i*i
	for i := 2; i*i <= n; i++ {
		if primes[i] {
			for j := i * i; j <= n; j += i {
				primes[j] = false
			}
		}
	}
	return primes
}

// eratosthenesSieveBitSet - за  O (N*loglogN)
// Со сжатием битов
func eratosthenesSieveBitSet(n int) bm_set.SetI {
	primes := bm_set.New(uint64(n + 1))
	for i := 0; i <= n; i++ {
		primes.Set(i)
	}
	primes.Delete(0)
	primes.Delete(1)
	// Для того чтобы найти все простые до n, достаточно выполнить просеивание только простыми, не превосходящими корня из n.
	// Так как все числа просты в изначальном созданном массиве, то все числа до SQRT(N) будут обработаны с последующим выкл/вкл до i*i
	for i := 2; i*i <= n; i++ {
		if primes.Get(i) {
			for j := i * i; j <= n; j += i {
				primes.Delete(j)
			}
		}
	}
	return primes
}

func TestEratosthenesSieve(t *testing.T) {
	primes := eratosthenesSieve(25)
	primesBs := eratosthenesSieveBitSet(25)
	assert.Equal(t, []bool{
		false,
		false,
		true,
		true,
		false,
		true,
		false,
		true,
		false,
		false,
		false,
		true,
		false,
		true,
		false,
		false,
		false,
		true,
		false,
		true,
		false,
		false,
		false,
		true,
		false,
		false,
	}, primes)

	for i := range primes {
		if primes[i] {
			if !primesBs.Get(i) {
				t.Errorf("true != set.Get() == ok, %d", i)
			}
		} else {
			if primesBs.Get(i) {
				t.Errorf("false != set.Get() == !ok, %d", i)
			}
		}
	}
}
