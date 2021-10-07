package graph

import (
	"fmt"
	"testing"
)

//var t = 0
//var enterTime []int
//var minTime []int
//var u []bool
//var g map[int][]int

func dfsB(v int, p int) {
	u[v] = true
	t++
	enterTime[v] = t
	minTime[v] = t
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
			dfsB(nv, v)
			if minTime[nv] < minTime[v] {
				minTime[v] = minTime[nv]
			}
			if minTime[nv] > enterTime[v] {
				// bridge
				fmt.Println(nv, v)
			}
		}

	}
}

// findBridges - find bridges in graph
// link to docs - https://e-maxx.ru/algo/bridge_searching
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
			dfsB(i, -1)
		}
	}
}

func Test_FindBridges(t *testing.T) {
	findBridges(5, [][]int{{1, 0}, {2, 0}, {3, 2}, {4, 2}, {4, 3}, {3, 0}, {4, 0}})
	// should print 1, 0
}
