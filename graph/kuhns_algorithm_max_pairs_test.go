package graph

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

// dfsKuhns: Пытается каждый раз увеличить число пар в паросочетании, true - если получилось создать новую пару
func dfsKuhns(used []bool, matched []int, gr map[int][]int, v int) bool {
	if used[v] {
		return false
	}

	used[v] = true

	nexts := gr[v]
	for i := range nexts {
		nx := nexts[i]
		// если вершина свободна (matched[nx] == -1), то можно сразу с ней соединиться
		// если она занята, то с ней можно соединиться только тогда,
		// когда из её текущей пары можно найти какую-нибудь другую
		// вершину для соединения + 1 (соединяем текущую вершину) -> увеличиваем кол-во пар
		if matched[nx] == -1 || dfsKuhns(used, matched, gr, matched[nx]) {
			matched[nx] = v
			return true
		}
	}
	return false
}

func cntPairs(n int, gr map[int][]int) int {
	cnt := 0
	used := make([]bool, n)
	matched := make([]int, n)
	for i := range matched {
		matched[i] = -1
	}

	for i := 0; i < n; i++ {
		for ui := range used {
			used[ui] = false
		}
		if dfsKuhns(used, matched, gr, i) {
			cnt++
		}
	}
	return cnt
}

func TestCntPairs(t *testing.T) {
	gr := map[int][]int{
		0: {1},
		1: {},
		2: {1, 3},
		3: {},
	}

	assert.Equal(t, cntPairs(4, gr), 2)
}
