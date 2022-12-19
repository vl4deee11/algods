package graph

import "fmt"

var t = 0
var enterTime []int
var minTime []int
var u []bool
var g map[int][]int

// dfsCutPoint - за время O(n+m)
// Пусть дан связный неориентированный граф.
// Точкой сочленения (или точкой артикуляции, англ. "cut vertex" или "articulation point") называется такая вершина,
// удаление которой делает граф несвязным.
func dfsCutPoint(v int, p int) {
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
	childrens := 0
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
			dfsCutPoint(nv, v)
			// Обновим minTime нашей текущей вершины v
			if minTime[nv] < minTime[v] {
				minTime[v] = minTime[nv]
			}
			// p != -1, тут сразу отсекаем корень дерева
			// Это условие главное для точки сочленения - звучит так, если есть ребенок у вершины у которого minTime[nv] >= enterTime[v],
			// то узел v является точкой сочленения, и не существует обратного ребра из nv в v
			if minTime[nv] >= enterTime[v] && p != -1 {
				// cutpoint
				fmt.Println(v)
			}
			childrens++
		}

	}
	// Рассмотрим теперь оставшийся случай: v = root.
	// Тогда эта вершина является точкой сочленения тогда и только тогда, когда эта вершина имеет более одного сына в дереве обхода в глубину.
	// В самом деле, это означает, что, пройдя из root по произвольному ребру, мы не смогли обойти весь граф,
	// откуда сразу следует, что root — точка сочленения.
	// Если бы мы прошли весть граф в глубину, то childrens == 1 или childrens == 0
	if p == -1 && childrens > 1 {
		// cutpoint
		fmt.Println(v)
	}
}

// findCutPoint - Ищет точки сочленения в графе
func findCutPoint(n int, connections [][]int) {
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
	// В отличие от поиска мостов, мы всегда идем в один узел
	dfsCutPoint(0, -1)
}
