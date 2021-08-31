package graph

import "fmt"

var t = 0
var enterTime []int
var minTime []int
var u []bool
var g map[int][]int

func dfsC(v int, p int) {
	u[v] = true
	t++
	enterTime[v] = t
	minTime[v] = t
	childrens := 0
	for i := range g[v] {
		nv := g[v][i]
		if nv == p {
			continue
		}

		if u[nv] {
			if enterTime[nv] < minTime[v] {
				minTime[v] = enterTime[nv]
			}
		} else {
			dfsC(nv, v)
			if minTime[nv] < minTime[v] {
				minTime[v] = minTime[nv]
			}
			if minTime[nv] >= enterTime[v] && p != -1 {
				// cutpoint
				fmt.Println(v)
			}
			childrens++
		}

	}
	if p == -1 && childrens > 1 {
		// cutpoint
		fmt.Println(v)
	}
}

// findCutPoint - find cutpoint in graph
// link to docs - https://e-maxx.ru/algo/cutpoints
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
	dfsC(0, -1)
}
