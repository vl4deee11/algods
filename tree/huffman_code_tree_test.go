package tree

import (
	"container/heap"
	"fmt"
	"testing"
	"unsafe"

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

// code - восстанавливает префикс код хаффмана, если текущий узел - правый ребенок родителя,
// устанавливаем бит в 1, иначе в ноль
func (n *node) code() (uint64, byte) {
	cp := n.p
	var (
		r    uint64
		bits byte
	)
	for cp != nil {
		if cp.r == n {
			r |= 1 << bits
		}
		bits++

		n = cp
		cp = cp.p
	}

	return r, bits
}

// decode: Декодирует строку из битов
func decode(de string, r *node) string {
	res := make([]byte, 0)
	cp := r
	for i := range de {
		if de[i] == '0' {
			cp = cp.l
		} else {
			cp = cp.r
		}
		if cp.b != 0 {
			res = append(res, cp.b)
			cp = r
		}
	}
	return *(*string)(unsafe.Pointer(&res))
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
//	                                         (c:1, b: 'd')     (c:2, b: 'b')
func Build(str string) (*node, map[byte]*node) {
	// Считаем частоты в строке
	hm := make(map[byte]int)
	for i := range str {
		hm[str[i]]++
	}

	b2n := make(map[byte]*node)
	leaves := make([]*node, 0, len(hm))
	for b := range hm {
		l := &node{c: hm[b], b: b}
		b2n[b] = l
		leaves = append(leaves, l)
	}

	// Инициализируем кучу для алгоритма
	h := MinHeap(leaves)
	heap.Init(&h)
	return buildTree(&h), b2n
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
	tr, b2n := Build(str)
	assert.Equal(t, tr.l.b, byte('a'))
	assert.Equal(t, tr.r.l.b, byte('c'))
	assert.Equal(t, tr.r.r.r.b, byte('b'))
	assert.Equal(t, tr.r.r.l.b, byte('d'))

	r, _ := b2n['a'].code()
	assert.Equal(t, fmt.Sprintf("%b", r), "0")

	r, _ = b2n['b'].code()
	assert.Equal(t, fmt.Sprintf("%b", r), "111")

	r, _ = b2n['c'].code()
	assert.Equal(t, fmt.Sprintf("%b", r), "10")

	r, _ = b2n['d'].code()
	assert.Equal(t, fmt.Sprintf("%b", r), "110")

	assert.Equal(t, decode("0000111111101010110", tr), "aaaabbcccd")
	assert.Equal(t, decode("1100", tr), "da")
}
