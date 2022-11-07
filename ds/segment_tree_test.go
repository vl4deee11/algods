package ds

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

// Дерево отрезков для операции сложения
// v - вершина
// arrTree - дерево, в 2 раза больше чем обычный массив, храним как 1 - корень, 2*r - левый ребенок, 2*(r+1) - правый ребенок
// [0, 1, 2, 3, 4, 5, 6, 7]
//            1
//       2    |    3
//   4      5 | 6     7
// tl, tr - ограничения для дерева отрезков - то есть то поддерево на котором работаем
func BuildSegTree(arr []int) []int {
	arrTree := make([]int, 4*len(arr))
	buildSegTree(arr, &arrTree, 1, 0, len(arr)-1)
	return arrTree
}

// buildSegTree Рекурсивно строит дерево
// Граничное условие когда приходим в лист (условие: tl == tr)
// записываем значение из массива в лист, при подъеме наверх складываем из узлов значения в корне
// Пример: [1, 2, 3, 4, 5, 6, 7, 8]
//                   36
//            10              26
//         3      7       11       15
//       1   2  3   4   5    6   7    8

func buildSegTree(arr []int, arrTree *[]int, v, tl, tr int) {
	if tl == tr {
		(*arrTree)[v] = arr[tl]
		return
	}
	tm := (tl + tr) / 2
	buildSegTree(arr, arrTree, left(v), tl, tm)
	buildSegTree(arr, arrTree, right(v), tm+1, tr)
	(*arrTree)[v] = (*arrTree)[left(v)] + /* или любая другая операция */ (*arrTree)[right(v)]
}

func left(i int) int {
	return 2 * i
}

func right(i int) int {
	return (2 * i) + 1
}

func parent(i int) int {
	return i / 2
}

func min(x, y int) int {
	if x < y {
		return x
	}
	return y
}

func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}

func OpT(arrTree []int, l, r int) int {
	return opT(arrTree, 1, 0, (len(arrTree)/4)-1, l, r)
}

// opT Выполняет операцию на дереве отрезков, на отрезке [l, r]
// Граничное условие l > r - не верно ушли по границам отрезка
// Граничное условие обрабатываемое поддерево на границах l == tl && r == tr
// В остальных случаях мы идем либо в правый отрезок и в правую ветвь либо в левый отрезок и в левую ветвь,
// или в обе ветви, при этом [tl, tr] являются границами для отрезка чья сумма лежит в arrTree[v]
// Так же наш отрезок [l, r] мы подрезаем, по границам обрабатываемого нами отрезка который лежит в arrTree[v]
// Одно из следствий l > r => при подрезке это условие станет верно и мы не уйдем в "неверную ветвь"
func opT(arrTree []int, v, tl, tr, l, r int) int {
	if l > r {
		return 0
	}
	if l == tl && r == tr {
		return arrTree[v]
	}
	tm := (tl + tr) / 2
	return opT(arrTree, left(v), tl, tm, l, min(r, tm)) + /* или любая другая операция */
		opT(arrTree, right(v), tm+1, tr, max(l, tm+1), r)
}

func Upd(arrTree *[]int, pos, newVal int) {
	upd(arrTree, 1, 0, (len(*arrTree)/4)-1, pos, newVal)
}

// upd Выполняет обновление на дереве отрезков
// Граничное условие tl == tr - попали в лист который надо обновить
// На каждом этапе принимаем решение куда идти в левый лист или в правый в зависимости от того, где обновляем значение
// При этом [tl, tr] являются границами для отрезка чья сумма лежит в arrTree[v]
// После обновления обновляем каждый раз корень
func upd(arrTree *[]int, v, tl, tr, pos, newVal int) {
	if tl == tr {
		(*arrTree)[v] = newVal
		return
	}
	tm := (tl + tr) / 2
	if pos <= tm {
		upd(arrTree, left(v), tl, tm, pos, newVal)
	} else {
		upd(arrTree, right(v), tm+1, tr, pos, newVal)
	}
	(*arrTree)[v] = (*arrTree)[left(v)] + /* или любая другая операция */
		(*arrTree)[right(v)]
}

func TestSegmentTree1(t *testing.T) {
	arr := []int{1, 2, 3, 4, 5, 6, 7, 8}
	//n := 8
	tree := BuildSegTree(arr)

	// 1. OpT 0 - n
	s1 := OpT(tree, 0, 4)
	assert.Equal(t, s1, 36)
	// 2. OpT 0 - 0
	s2 := OpT(tree, 0, 0)
	assert.Equal(t, s2, 1)
	// 3. OpT 2 - 4
	s3 := OpT(tree, 2, 4)
	assert.Equal(t, s3, 12)

	// 4. Upd 3->11 & OpT 2 - 4
	Upd(&tree, 2, 11)
	s4 := OpT(tree, 2, 4)
	assert.Equal(t, s4, 20)
}

func TestSegmentTree2(t *testing.T) {
	arr := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}
	n := 9
	tree := BuildSegTree(arr)

	// 1. OpT 0 - n
	s1 := OpT(tree, 0, n)
	assert.Equal(t, s1, 45)
	// 2. OpT 0 - 0
	s2 := OpT(tree, 0, 0)
	assert.Equal(t, s2, 1)
	// 3. OpT 2 - 4
	s3 := OpT(tree, 2, 4)
	assert.Equal(t, s3, 12)

	// 4. Upd 3->11 & OpT 2 - 4
	Upd(&tree, 2, 11)
	s4 := OpT(tree, 2, 4)
	assert.Equal(t, s4, 20)
}
