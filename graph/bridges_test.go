package graph

import (
	"fmt"
	"testing"
)

// var t = 0
// var enterTime []int
// var minTime []int
// var u []bool
// var g map[int][]int

// dfsBridges - за время O(n+m)
// Пусть дан неориентированный граф.
// Мостом называется такое ребро, удаление которого делает граф несвязным (или, точнее, увеличивает число компонент связности).
// Требуется найти все мосты в заданном графе.
func dfsBridges(v int, p int) {
	// Ставим метку того, что уже посещали вершину
	u[v] = true
	// Увеличиваем время захода при поиске
	t++
	// Время входа в вершину
	enterTime[v] = t
	// Время minTime[v] равно минимуму из
	// enterTime[v] - времени захода в саму вершину,
	// enterTime[p] - времён захода в каждую вершину p, являющуюся концом некоторого обратного ребра (v, p),
	// minTime[to]  - времён для каждой вершины to, являющейся непосредственным сыном v в дереве поиска
	// minTime - является рекурсивным минимум
	minTime[v] = t
	for i := range g[v] {
		nv := g[v][i]
		if nv == p {
			// Это критерий прохода по ребру дерева поиска в обратную сторону, мы его не обрабатываем
			continue
		}

		if u[nv] {
			// Если вершину nv мы уже обрабатывали ранее, то обновим minTime нашей текущей вершины v,
			// таким образом мы словим обратное ребро
			if enterTime[nv] < minTime[v] {
				minTime[v] = enterTime[nv]
			}
		} else {
			// Иначе, вершину nv мы еще не обрабатывали
			dfsBridges(nv, v)
			// Обновим minTime нашей текущей вершины v
			if minTime[nv] < minTime[v] {
				minTime[v] = minTime[nv]
			}
			// Это условие главное для моста - звучит так, если есть ребенок у вершины у которого minTime[nv] > enterTime[v],
			// то ребро nv - v является мостом, и не существует обратного ребра из nv в v
			// Ребро (v, nv) является мостом тогда и только тогда, когда (v,nv) ∈ T и из вершины nv и любого ее потомка нет
			// обратного ребра в вершину v или предка v
			if minTime[nv] > enterTime[v] {
				// bridge
				fmt.Println(nv, v)
			}
		}

	}
}

// findBridges - Ищет мосты в графе
func findBridges(n int, connections [][]int) {
	t = 0
	enterTime = make([]int, n)
	minTime = make([]int, n)
	u = make([]bool, n)
	g = make(map[int][]int)
	for i := range connections {
		g[connections[i][0]] = append(g[connections[i][0]], connections[i][1])
		g[connections[i][1]] = append(g[connections[i][1]], connections[i][0])
	}
	for i := 0; i < n; i++ {
		u[i] = false
	}
	for i := 0; i < n; i++ {
		if !u[i] {
			dfsBridges(i, -1)
		}
	}
}

func Test_FindBridges(t *testing.T) {
	findBridges(5, [][]int{{1, 0}, {2, 0}, {3, 2}, {4, 2}, {4, 3}, {3, 0}, {4, 0}})
	// should print 1, 0
}
