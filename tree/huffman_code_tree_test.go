package tree

import (
	"container/heap"
	"testing"

	"github.com/stretchr/testify/assert"
)

type node struct {
	// p родитель узла
	p *node
	// l левый ребенок узла
	l *node
	// r правый ребенок узла
	r *node
	// c частота появления символа в тексе
	c int
	// b байт символа в который в листе
	b byte
}

// Куча для построения дерева Хаффмана
type MinHeap []*node

func (h MinHeap) Len() int           { return len(h) }
func (h MinHeap) Less(i, j int) bool { return h[i].c < h[j].c }
func (h MinHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }

func (h *MinHeap) Push(x interface{}) {
	*h = append(*h, x.(*node))
}

func (h *MinHeap) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}

// Пример дерева для строки aaaabbcccd
//
//	                     (c: 10)
//	(c:4, b: 'a')                         (c:6)
//	                       (c:3, b: 'c')                  (c:3)
//	                                         (c:2, b: 'b')     (c:1, b: 'd')
func Build(str string) *node {
	// Считаем частоты в строке
	hm := make(map[byte]int)
	for i := range str {
		hm[str[i]]++
	}

	leaves := make([]*node, 0, len(hm))
	for b := range hm {
		leaves = append(leaves, &node{c: hm[b], b: b})
	}

	// Инициализируем кучу для алгоритма
	h := MinHeap(leaves)
	heap.Init(&h)
	return buildTree(&h)
}

// buildTree - строит дерево хаффмана.
// Используется жадный алгоритм на минимальной куче, жадность заключается в том, что чем меньше частота,
// тем раньше будет вытянут узел и тем ниже он окажется в дереве (то есть до него будет длиннее префиксный код), и наоборот
func buildTree(h *MinHeap) *node {
	if h.Len() == 0 {
		return nil
	}

	var l, r *node
	for h.Len() > 1 {
		l, r = heap.Pop(h).(*node), heap.Pop(h).(*node)
		p := &node{l: l, r: r, c: l.c + r.c}
		l.p = p
		r.p = p
		heap.Push(h, p)
	}

	return (*h)[0]
}

func TestHuffmanCodeTree(t *testing.T) {
	str := "aaaabbcccd"
	tr := Build(str)
	assert.Equal(t, tr.l.b, byte('a'))
	assert.Equal(t, tr.r.l.b, byte('c'))
	assert.Equal(t, tr.r.r.r.b, byte('b'))
	assert.Equal(t, tr.r.r.l.b, byte('d'))
}
