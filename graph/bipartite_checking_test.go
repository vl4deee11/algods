package graph

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// Пусть дан неориентированный граф. Требуется проверить, является ли он двудольным,
// т.е. можно ли разделить его вершины на две доли так, чтобы не было рёбер, соединяющих две вершины одной доли.
// Если граф является двудольным, то вывести сами доли.

// bipartiteChecking - за O(M)
// Здесь n - кол-во вершин в графе, st - стартовая вершина (откуда искать),
// edges - ребро, такое что edges[i][0] - вершина 1, edges[i][1] - вершина 2 (ребра не направленные)
func bipartiteChecking(n int, edges [][2]int) ([]int, bool) {
	graph := make(map[int][]int, n)
	for _, edge := range edges {
		u, v := edge[0], edge[1]
		graph[u] = append(graph[u], v)
		graph[v] = append(graph[v], u)
	}

	part := make([]int, n)
	for i := range part {
		// Все вершины пока не посещены и не находятся в долях
		part[i] = -1
	}

	// Признак двудольности
	// Теорема. Граф является двудольным тогда и только тогда, когда все его простые циклы имеют чётную длину.
	// С практической точки зрения искать все простые циклы неудобно. Намного проще проверять граф на двудольность следующим алгоритмом:

	// Произведём серию поисков в ширину.
	// Т.е. будем запускать поиск в ширину из каждой не посещённой вершины.
	// Ту вершину, из которой мы начинаем идти, мы помещаем в первую долю. В процессе поиска в ширину,
	// если мы идём в какую-то новую вершину, то мы помещаем её в долю, отличную от доли текущей вершину.
	// Если же мы пытаемся пройти по ребру в вершину, которая уже посещена, то мы проверяем,
	// чтобы эта вершина и текущая вершина находились в разных долях. В противном случае граф двудольным не является.
	// По окончании работы алгоритма мы либо обнаружим, что граф не двудолен, либо найдём разбиение вершин графа на две доли.

	// Дополнение - если существует связь из u - v,
	// то мы либо покрасим вершины в разные цвета,
	// либо проверим вершину v или u (в зависимости от того откуда начали) и если проверка не прошла,
	// значит связь уже была построена ранее и построить разбиение графа попросту невозможно
	ok := true
	for i := 0; i < n; i++ {
		if part[i] != -1 {
			// Уже посетили ранее
			continue
		}
		q := newQueue()
		q.Enqueue(i)
		// Кладем в группу 0
		part[i] = 0
		for q.first != nil {
			cn := q.Dequeue().(int)
			conns := graph[cn]
			for ni := range conns {
				nx := conns[ni]
				if part[nx] == -1 {
					// Кладем в отличную группу
					if part[cn] == 0 {
						part[nx] = 1
					} else {
						part[nx] = 0
					}
					q.Enqueue(nx)
				} else {
					// Проверяем что в разных группах
					ok = part[nx] != part[cn]
				}
			}
		}
	}
	return part, ok
}

// Краткая очередь для БФС
type queue struct {
	first *node
	last  *node
}

type node struct {
	next *node
	data interface{}
}

func newQueue() *queue {
	return new(queue)
}

func (eq *queue) Enqueue(e interface{}) {
	if eq.last != nil {
		eq.last.next = new(node)
		eq.last.next.data = e
		eq.last = eq.last.next
		return
	}

	ln := &node{data: e}
	eq.last = ln
	eq.first = ln
}

func (eq *queue) Dequeue() (e interface{}) {
	if eq.first == nil {
		return nil
	}
	v := eq.first
	eq.first = eq.first.next
	if eq.first == nil {
		eq.last = nil
	}
	return v.data
}

func TestBipartiteChecking(t *testing.T) {
	n := 6
	// 0 --- 3
	//       |
	// 1 ----|
	// 4--------2
	//
	// 5
	part, ok := bipartiteChecking(n, [][2]int{{0, 3}, {1, 3}, {4, 2}})
	assert.Equal(t, true, ok)
	assert.Equal(t, []int{0, 0, 0, 1, 1, 0}, part)

	// 0 --- 3--|
	//       |  |
	// 1 ----|  |
	// |--------2
	// 4
	//
	// 5
	part, ok = bipartiteChecking(n, [][2]int{{0, 3}, {1, 3}, {2, 3}, {2, 1}})
	assert.Equal(t, false, ok)
}
