package graph

import (
	"sort"
	"testing"

	"github.com/stretchr/testify/assert"
)

// kruskalMST - за O(M log N)
// Дан взвешенный неориентированный граф. Требуется найти такое поддерево этого графа, которое бы соединяло все его вершины,
// и при этом обладало наименьшим весом (т.е. суммой весов рёбер) из всех возможных.
// Такое поддерево называется минимальным остовным деревом или простом минимальным остовом (minimum spanning tree - MST).
//
// Алгоритм Крускала изначально помещает каждую вершину в своё дерево, а затем постепенно объединяет эти деревья,
// объединяя на каждой итерации два некоторых дерева некоторым ребром. Перед началом выполнения алгоритма,
// все рёбра сортируются по весу (в порядке неубывания).
// Затем начинается процесс объединения: перебираются все рёбра от первого до последнего (в порядке сортировки),
// и если у текущего ребра его концы принадлежат разным поддеревьям, то эти поддеревья объединяются,
// а ребро добавляется к ответу. По окончании перебора всех рёбер все вершины окажутся принадлежащими одному поддереву, и ответ найден.
// Здесь n - кол-во вершин в графе, edges - ребро, такое что edges[i][0] - вершина 1, edges[i][1] - вершина 2, edges[i][2] - вес
func kruskalMST(n int, edges [][3]int) [][3]int {
	uff := new(ufmst)
	uff.r = make([]int, n)
	uff.p = make([]int, n)
	// Инициализация Union Find Minimal Spanning Tree, сначала каждый узел лежит в свое поддереве
	for i := 0; i < n; i++ {
		uff.p[i] = i
	}

	// Сортируем ребра в порядке возрастания
	sort.Slice(edges, func(i, j int) bool {
		return edges[i][2] < edges[j][2]
	})

	// Результирующие ребра для минимального остовного дерева
	res := make([][3]int, 0)

	// Тут происходит жадная стратегия, начиная с самых легких ребер, начинаем объединять узлы если они еще не в одном множестве
	for i := 0; i < len(edges); i++ {
		a, b := edges[i][0], edges[i][1]
		if uff.find(a) != uff.find(b) {
			uff.union(a, b)
			res = append(res, edges[i])
		}
	}
	return res
}

// ufmst - Union Find Minimal Spanning Tree - краткая реализация непересекающегося множества для алгоритма kruskalMST
type ufmst struct {
	p, r []int
}

func (uf *ufmst) find(e int) int {
	if uf.p[e] == e {
		return e
	}
	uf.p[e] = uf.find(uf.p[e])
	return uf.p[e]
}

func (uf *ufmst) union(e1, e2 int) {
	r1 := uf.find(e1)
	r2 := uf.find(e2)
	if r1 == r2 {
		return
	}

	switch {
	case uf.r[r1] < uf.r[r2]:
		uf.p[r1] = r2
	case uf.r[r1] > uf.r[r2]:
		uf.p[r2] = r1
	default:
		uf.p[r2] = r1
		uf.r[r1]++
	}
	return
}

func TestKruskalMST(t *testing.T) {
	// Тестовый пример взял из https://upload.wikimedia.org/wikipedia/commons/0/01/MST_Kruskal.gif
	edges := [][3]int{{0, 4, 1}, {0, 1, 3}, {1, 2, 5}, {1, 4, 4}, {2, 3, 2}, {2, 4, 6}, {3, 4, 7}}
	res := kruskalMST(5, edges)
	// Сумма весов = 11, можно показать что легче уже не выйдет создать дерева
	assert.Equal(t, [][3]int{{0, 4, 1}, {2, 3, 2}, {0, 1, 3}, {1, 2, 5}}, res)
}
