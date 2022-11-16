package ds

import (
	"math"
	"sort"
	"testing"

	"github.com/stretchr/testify/assert"
)

// Алгоритм Мо — применяется для решения задач, в которых требуется отвечать на запросы arr[L…R]
// на массиве без изменения элементов в оффлайн за время O(Q*logQ + (N+Q)*SQRT(N)),
// где Q — количество запросов, а N — количество элементов в массиве.
// Характерными примерами задач на этот алгоритм являются: нахождение моды на отрезке (число, которое встречается больше всех остальных),
// вычисление количества инверсий на отрезке.

// ====================================================================================
// В данном примере: мы считаем операции количество различных чисел на отрезке [Li, Ri]
// В данном алгоритме меняются только функции Add, Del
// ====================================================================================

// Сгруппируем все запросы в блоки размера С = SQRT(N)
// по их левой границе и внутри каждого блока отсортируем запросы по правой границе

// Q: структура запроса, [L,R] - границы запроса; idx - индекс (или номер) запроса
type Q struct {
	L, R, idx int
}

func getc(arr []int) int {
	return int(math.Sqrt(float64(len(arr))) + 1)
}

// BuildBuckets: инициализирует структуру бакетов
func BuildBuckets(arr []int, queries []*Q) [][]*Q {
	// Кол-во бакетов
	c := getc(arr)

	buckets := make([][]*Q, c)
	// группировка запросов по началу отрезков Q.L
	for i := range queries {
		if queries[i].L >= len(arr) || queries[i].L < 0 {
			continue
		}
		if queries[i].R >= len(arr) || queries[i].R < 0 {
			continue
		}
		buckets[queries[i].L/c] = append(buckets[queries[i].L/c], queries[i])
	}

	// сортировка каждого бакета
	for i := range buckets {
		sort.Slice(buckets[i], func(ki, kj int) bool {
			return buckets[i][ki].R < buckets[i][kj].R
		})
	}

	return buckets
}

// res: равная числу различных элементов на текущем отрезков,
// то есть числу ненулевых элементов массива cnt (изначально ноль)

// В данном примере: мы считаем операции количество различных чисел на отрезке [Li, Ri]
// Add: добавляем arr[k] на текущий отрезок
func Add(arr []int, cnt *[]int, k int, res *int) {
	//  Или другая операция
	if (*cnt)[arr[k]] == 0 {
		*res++
	}
	(*cnt)[arr[k]]++
}

// Add: удаляем arr[k] с текущего отрезка
func Del(arr []int, cnt *[]int, k int, res *int) {
	//  Или другая операция
	(*cnt)[arr[k]]--
	if (*cnt)[arr[k]] == 0 {
		*res--
	}
}

// ProcessQ: Выполняет все запросы на бакетах
// buckets - бакеты из BuildBuckets
// arr - изначальный массив чисел
// lq - длинна массива ответов
// SQRT(N) - кол-во блоков
// Q - Кол-во запросов
// N - длинна массива
// Трюк в том, что правая граница суммарно сдвинется на O(N) (по одному блоку), потому что отрезки отсортированы,
// а левая каждый раз сдвинется не более чем на O(SQRT(N)), так как все левые границы сгруппированы по корневым блокам.
// Изменение левых границ суммарно по всем блокам займёт O(Q*SQRT(N)) операций, а правых O(N*SQRT(N)),
// так что итоговая асимптотика решения будет O(Q*SQRT(N) + N*SQRT(n)).
func ProcessQ(buckets [][]*Q, arr []int, lq int) []int {
	// Кол-во бакетов
	c := getc(arr)

	ans := make([]int, lq)

	// cnt: массив размера N,
	// в котором для каждого значения хранится число элементов
	// на текущем отрезке с данным значением (изначально заполнен нулями);
	// [ТОЛЬКО: для конкретной задачи]
	cnt := make([]int, len(arr))

	for i := 0; i < c; i++ {

		// Обнуляем границы
		l := i * c
		r := i*c - 1
		// Обнуляем счетчики [ТОЛЬКО: для конкретной задачи]
		for ci := range cnt {
			cnt[ci] = 0
		}

		res := 0
		for _, q := range buckets[i] {
			// Двигаем правую границу, пока правая граница не дошла до границы запроса
			for r < q.R {
				r++
				Add(arr, &cnt, r, &res)
			}
			// Дальше делаем так, чтобы левая граница совпала
			for l < q.L {
				Del(arr, &cnt, l, &res)
				l++
			}
			for l > q.L {
				l--
				Add(arr, &cnt, l, &res)
			}
			ans[q.idx] = res
		}
	}
	return ans
}

func TestAlgoMO(t *testing.T) {
	arr := []int{1, 2, 3, 3, 3, 4, 5, 6}
	queries := []*Q{
		{
			L:   0,
			R:   len(arr) - 1,
			idx: 0,
		},
		{
			L:   0,
			R:   0,
			idx: 1,
		},
		{
			L:   0,
			R:   4,
			idx: 2,
		},
		{
			L:   len(arr) - 1,
			R:   len(arr) - 1,
			idx: 3,
		},
		{
			L:   4,
			R:   5,
			idx: 4,
		},
		{
			L:   3,
			R:   5,
			idx: 5,
		},
	}
	buckets := BuildBuckets(arr, queries)

	answers := ProcessQ(buckets, arr, len(queries))
	assert.Equal(t, answers[0], 6)
	assert.Equal(t, answers[1], 1)
	assert.Equal(t, answers[2], 3)
	assert.Equal(t, answers[3], 1)
	assert.Equal(t, answers[4], 2)
	assert.Equal(t, answers[5], 2)

	arr = []int{1, 2, 3, 6, 3, 4, 5}
	queries = []*Q{
		{
			L:   0,
			R:   len(arr) - 1,
			idx: 0,
		},
		{
			L:   0,
			R:   0,
			idx: 1,
		},
		{
			L:   0,
			R:   4,
			idx: 2,
		},
		{
			L:   len(arr) - 1,
			R:   len(arr) - 1,
			idx: 3,
		},
		{
			L:   4,
			R:   5,
			idx: 4,
		},
		{
			L:   3,
			R:   5,
			idx: 5,
		},
	}
	buckets = BuildBuckets(arr, queries)

	answers = ProcessQ(buckets, arr, len(queries))
	assert.Equal(t, answers[0], 6)
	assert.Equal(t, answers[1], 1)
	assert.Equal(t, answers[2], 4)
	assert.Equal(t, answers[3], 1)
	assert.Equal(t, answers[4], 2)
	assert.Equal(t, answers[5], 3)
}
