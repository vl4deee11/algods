package combinatorial

import (
	"fmt"
	"math"
	"math/rand"
	"testing"
	"time"
)

// Алгоритм имитации отжига — эвристический алгоритм глобальной оптимизации,
// особенно эффективный при решении дискретных и комбинаторных задач.
// Общее описание
// Пусть имеется некоторая функция f(x) от состояния, которую мы хотим минимизировать.
// Возьмём в качестве базового решения какое-то состояние x_0 и будем пытаться его улучшать
// Введём температуру t — какое-то действительное число, изначально равное единице, которое будет изменяться в течение оптимизации
// и влиять на вероятность перейти в соседнее состояние.
// Пока не придём к оптимальному решению или пока не закончится время (k - кол-во шагов),
// будем повторять следующие шаги:
// 	1. Уменьшим температуру t_i = T(t_i-1) по какой-то формуле T
//  2. Выберем случайного соседа x: какое-то состояние y, которое может быть получено из x каким-то небольшим изменением.
//  3. С вероятностью P(f(x), f(y), t_i) сделаем присвоение x <- y, иначе оставим x как есть

// Рассмотрим его на примере задачи
// Дана шахматная доска размером n * n, и n ферзей. Нужно расставить их так, чтобы они не били друг друга.
// Выберем функцию f(p, n), равную числу успешно расставленных ферзей. Это и будет функция на которой мы будем отжигать
//

func abs(x int) int {
	if x >= 0 {
		return x
	}
	return -x
}

// Кодируем как i ферзь стоит на i строке и на p[i] столбце, т.е координаты (i, p[i])
func f(p []int, n int) int {
	s := 0
	for i := 0; i < n; i++ {
		d := 1
		for j := 0; j < i; j++ {
			// Если не бьем по диагонали никого и предыдущих
			if abs(i-j) == abs(p[i]-p[j]) {
				d = 0
			}
		}
		s += d
	}
	return s
}

func shuffle(a []int) {
	rand.Shuffle(len(a), func(i, j int) { a[i], a[j] = a[j], a[i] })
}

// tf - формула по которой мы уменьшаем температуру
func tf(ti float64) float64 {
	return ti * 0.99
}

func annealing(n, k int) int {
	rand.Seed(time.Now().UnixNano())
	// k - кол-во итераций алгоритма.
	// Генерируем начальную перестановку
	var p = make([]int, n)
	for i := range p {
		p[i] = i
	}
	// Перемешиваем изначальную перестановку
	shuffle(p)
	// Пока что это лучший ответ
	ans := f(p, n)

	// Далее начинаем отжигать
	t := 1.0
	for i := 0; i < k && ans < n; i++ {
		t = tf(t)
		u := make([]int, n)
		copy(u, p)
		i1 := rand.Intn(n)
		i2 := rand.Intn(n)
		u[i1], u[i2] = u[i1], u[i2]
		val := f(u, n)
		if val > ans || rand.Float64() < math.Exp(float64(val-ans)/t) {
			p[i1], p[i2] = p[i1], p[i2]
			ans = val
		}
	}

	return ans
}

func TestAnnealing(t *testing.T) {
	ans := annealing(100, 1000)
	fmt.Println(ans)

}
