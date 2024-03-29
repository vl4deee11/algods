package graph

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// Дан ориентированный или неориентированный взвешенный граф G с n вершинами. Требуется найти значения всех величин d(i, j) — длины кратчайшего пути из вершины i в вершину j.
//
// Предполагается, что граф не содержит циклов отрицательного веса,
// (тогда ответа между некоторыми парами вершин может просто не существовать — он будет бесконечно маленьким).
//  Замечание: Данный алгоритм можно применять для всех ребер и даже для ребер с весом меньше нуля. А если в графе таких весов не будет,
// можно посчитать через Дейкстру по всем вершинам, асимптотически это будет O(N^2*logN) - что быстрее чем O (N^3).

// floydWarshall - за  O (N^3).
// Здесь n - кол-во вершин в графе, st - стартовая вершина (откуда искать),
// edges - ребро, такое что edges[i][0] - вершина 1, edges[i][1] - вершина 2, edges[i][2] - вес (ребра не направленные)
func floydWarshall(n int, edges [][3]int) [][]int {
	// Матрица расстояний
	dist := make([][]int, n)
	inf := 1 << 31

	for i := range dist {
		dist[i] = make([]int, n)
		for j := range dist[i] {
			if i == j {
				dist[i][j] = 0
			} else {
				dist[i][j] = inf
			}
		}
	}

	// Поставляем связи
	for i := range edges {
		dist[edges[i][0]][edges[i][1]] = edges[i][2]
	}

	// Ключевая идея алгоритма — разбиение процесса поиска кратчайших путей на фазы.
	//
	// Перед k-ой фазой (k = 0...n-1) считается, что в матрице расстояний d[][] сохранены длины таких кратчайших путей,
	// которые содержат в качестве внутренних вершин только вершины из множества { 0, 1, 2, ..., k-1 }.
	// Иными словами, перед k-ой фазой величина d[i][j] равна длине кратчайшего пути из вершины i в вершину j,
	// если этому пути разрешается заходить только в вершины с номерами, меньшими k (начало и конец пути не считаются).
	// Пусть теперь мы находимся на k-ой фазе, и хотим пересчитать матрицу d[][] таким образом, чтобы она соответствовала требованиям уже для k+1-ой фазы. Зафиксируем какие-то вершины i и j. У нас возникает два принципиально разных случая:
	//
	// 1. Кратчайший путь из вершины i в вершину j, которому разрешено дополнительно проходить через вершины { 0, 1, 2, ..., k },
	// совпадает с кратчайшим путём, которому разрешено проходить через вершины множества { 0, 1, 2, ..., k-1 }.
	// В этом случае величина d[i][j] не изменится при переходе с k-ой на k+1-ую фазу.
	//
	// 2. "Новый" кратчайший путь стал лучше "старого" пути.
	// Это означает, что "новый" кратчайший путь проходит через вершину k.
	// Сразу отметим, что мы не потеряем общности, рассматривая далее только простые пути (т.е. пути, не проходящие по какой-то вершине дважды).
	//
	// Тогда заметим, что если мы разобьём этот "новый" путь вершиной k на две половинки (одна идущая i => k, а другая — k => j),
	// то каждая из этих половинок уже не заходит в вершину k.
	// Но тогда получается, что длина каждой из этих половинок была посчитана ещё на k-1-ой фазе или ещё раньше,
	// и нам достаточно взять просто сумму d[i][k] + d[k][j], она и даст длину "нового" кратчайшего пути.
	for k := 0; k < n; k++ {
		for i := 0; i < n; i++ {
			for j := 0; j < n; j++ {
				// Пример для понимания: допустим нет прямого пути между 1 и 0 вершиной, а между ними есть такая связь
				// 0 -> 2 -> 3 -> 1, то на k = 2 найдем путь 0 -> 3, на k = 3 найдем путь 0 -> 1 = [0 -> 3, 3 -> 1] и тд
				if dist[i][j] > dist[i][k]+dist[k][j] {
					dist[i][j] = dist[i][k] + dist[k][j]
				}
			}
		}
	}
	return dist
}

func TestFloydWarshall(t *testing.T) {
	// Тестовый пример взял из https://upload.wikimedia.org/wikipedia/commons/thumb/2/2e/Floyd-Warshall_example.svg/2880px-Floyd-Warshall_example.svg.png
	n := 4

	edges := [][3]int{
		{1, 0, 4},
		{0, 2, -2},
		{1, 2, 3},
		{2, 3, 2},
		{3, 1, -1},
	}
	d := floydWarshall(n, edges)
	assert.Equal(t, [][]int{
		{0, -1, -2, 0},
		{4, 0, 2, 4},
		{5, 1, 0, 2},
		{3, -1, 1, 0},
	}, d)
}
