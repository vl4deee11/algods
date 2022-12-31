package graph

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

var color []int
var cycleSt, cycleEnd int

func dfsCycleCheck(gr map[int][]int, v int) bool {
	// Красим вершину в цвет1
	color[v] = 1

	nxt := gr[v]
	for i := range nxt {
		nx := nxt[i]
		// Если есть сосед который мы не посещали == цвет0, то пойдем в него
		if color[nx] == 0 {
			// Нашли цикл где-то в глубине - выходим сразу
			if dfsCycleCheck(gr, nx) {
				return true
			}
		} else if color[nx] == 1 {
			// Если мы нашли вершину с цвет1, то это значит что, при поиске в глубины из вершины nx мы еще не вышли,
			// но теперь вторично можем в нее попасть => цикл существует.
			// Если граф неориентированный, то случаи, когда поиск в глубину из какой-то вершины пытается пойти в предка, не считаются
			cycleSt = nx
			cycleEnd = v
			return true
		}
	}
	// При выходе из вершины красим в цвет2
	color[v] = 2
	return false
}

// cycleCheck - за O(M)
// Здесь n - кол-во вершин в графе,
// edges - ребро, такое что edges[i][0] - вершина 1, edges[i][1] - вершина 2 (ребра направленные, но данный алгоритм работает и в не направленных ребрах)
func cycleCheck(n int, edges [][2]int) (int, int, bool) {
	// Начало и конец цикла
	cycleSt = -1
	cycleEnd = -1
	// Цвета вершин
	color = make([]int, n)
	// Обычный граф
	gr := make(map[int][]int)

	for i := range edges {
		gr[edges[i][0]] = append(gr[edges[i][0]], edges[i][1])
	}

	for i := 0; i < n; i++ {
		// Произведем серию поисков в глубину, если есть вершина которую мы еще не посещали, мы пойдем в нее
		if dfsCycleCheck(gr, i) {
			return cycleSt, cycleEnd, true
		}
	}
	return -1, -1, false
}

func TestCycleCheck(t *testing.T) {
	n := 5
	// 0 -> 1  <- 2
	//      |     ^
	//      v     |
	//      3  -> 4
	st, end, ok := cycleCheck(n, [][2]int{{0, 1}, {2, 1}, {1, 3}, {3, 4}, {4, 2}})
	assert.Equal(t, true, ok)
	assert.Equal(t, 1, st)
	assert.Equal(t, 2, end)

	// 0 -> 1  -> 2
	//      |     ^
	//      v     |
	//      3  -> 4
	st, end, ok = cycleCheck(n, [][2]int{{0, 1}, {1, 2}, {1, 3}, {3, 4}, {4, 2}})
	assert.Equal(t, false, ok)
}
