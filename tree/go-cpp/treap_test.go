package tree

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// Декартово дерево или диамида (англ. treap) — это структура данных,
// объединяющая в себе бинарное дерево поиска и бинарную кучу
// Более строго, это структура данных, которая хранит пары (X,Y) в виде бинарного дерева таким образом, что она является
// бинарным деревом поиска по x и бинарной пирамидой по y. Предполагая, что все X и все Y являются различными, получаем,
// что если некоторый элемент дерева содержит (X0,Y0), то у всех элементов в левом поддереве X < X0,
// у всех элементов в правом поддереве X > X0, а также и в левом, и в правом поддереве имеем: Y < Y0.
type treapNode struct {
	// l, r - левый и правые узлы корня соответственно
	l, r *treapNode
	// Y - prior - приоритет дерева
	// X - key - ключ дерева
	prior, key int
}

// search - за O(log N) в среднем
func search(root *treapNode, x int) *treapNode {
	// Ищет элемент с указанным значением ключа x. Реализуется абсолютно так же, как и для обычного бинарного дерева поиска.
	if root.key == x {
		return root
	}
	if root.key > x {
		return search(root.l, x)
	}
	return search(root.r, x)
}

// split - за O(log N).
// Разделяет дерево T на два дерева L и R (которые являются возвращаемым значением) таким образом,
// что L содержит все элементы, меньшие по ключу X, а R содержит все элементы, большие X
func split(root *treapNode, l, r **treapNode, x int) {
	if root == nil {
		*l = nil
		*r = nil
		return
	}
	if root.key > x {
		// И дальше рекурсивно идем в лево, при этом левый указатель оставляем как есть,
		// а правый указатель заменяем на r.l, так как все <= чем root.key, но больше чем x, должно быть в левом поддереве r
		// а правое поддерево root будет без изменений
		split(root.l, l, &root.l, x)
		// Подвешиваем корень за правое дерево
		*r = root
	} else {
		// И дальше рекурсивно идем в право, при этом правый указатель оставляем как есть,
		// а левый указатель заменяем на l.r, так как все >= чем root.key, но меньше чем x, должно быть в правом поддереве l
		// а левое поддерево root будет без изменений
		split(root.r, &root.r, r, x)
		// Подвешиваем корень за левое дерево
		*l = root
	}
}

// merge - за O(log N).
// Объединяет два поддерева L и R, и возвращает это новое дерево.
// Она работает в предположении, что L и R обладают соответствующим порядком (все значения X в первом [L] меньше значений X во втором [R]).
// Таким образом, нам нужно объединить их так, чтобы не нарушить порядок по приоритетам Y.
// Для этого просто выбираем в качестве корня то дерево, у которого Y в корне больше, и рекурсивно вызываем себя от
// другого дерева и соответствующего сына выбранного дерева.
func merge(root **treapNode, l, r *treapNode) {
	if l == nil {
		*root = r
		return
	} else if r == nil {
		*root = l
		return
	}

	if l.prior > r.prior {
		// Кладем в корень тот узел, где выше приоритет (не нарушим кучу по Y)
		*root = l

		// Зная, что по условию бин дерева в l.l все ключи <= ключу в корне
		// И по условию пирамиды все приоритеты в l.l <= приоритету в корне l => левая часть root уже правильная
		// и надо достроить только правую часть и именно поэтому мы идем в правый узел root -> так же переключаем в правый узел l
		merge(&(*root).r, l.r, r)
	} else {
		// Кладем в корень тот узел, где выше приоритет (не нарушим кучу по Y)
		*root = r

		// Зная, что по условию бин дерева в r.r все ключи >= ключу в корне
		// И по условию пирамиды все приоритеты в r.r <= приоритету в корне r => правая часть root уже правильная
		// и надо достроить только левую часть и именно поэтому мы идем в левый узел root -> так же переключаем в левый узел r
		merge(&(*root).l, l, r.l)
	}
}

// insert - за O(log N)
// Реализация Insert (X, Y).
// Сначала спускаемся по дереву (как в обычном бинарном дереве поиска по X),
// но останавливаемся на первом элементе, в котором значение приоритета
// оказалось меньше Y. Мы нашли позицию, куда будем вставлять наш элемент.
// Теперь вызываем Split (X) от найденного элемента (от элемента вместе со всем его поддеревом),
// и возвращаемые ею L и R записываем в качестве левого и правого сына добавляемого элемента.
func insert(root **treapNode, it *treapNode) {
	if *root == nil {
		*root = it
		return
	}
	if it.prior > (*root).prior {

		// Здесь вызов означает вызов функции: "разделите treap root по значению it.key
		// на два treap и сохраните левые treap it.l и правый treap it.r",
		// Разделив по ключу it.key => мы поддерживаем свойства бинарного дерева по it.key
		// и свойства бинарной кучи по it.prior
		split(*root, &it.l, &it.r, it.key)
		*root = it
	} else {
		// Спуск как в обычно бинарном дереве по X (it.key)
		if (*root).key <= it.key {
			insert(&(*root).r, it)
		} else {
			insert(&(*root).l, it)
		}
	}
}

// erase - за O(log N)
// Реализация Erase (X).
// Спускаемся по дереву (как в обычном бинарном дереве поиска по X), ища удаляемый элемент.
// Найдя элемент, мы просто вызываем Merge от его левого и правого сыновей,
// и возвращаемое ею значение ставим на место удаляемого элемента.
func erase(root **treapNode, key int) {
	if (*root).key == key {
		merge(root, (*root).l, (*root).r)
		return
	}
	if (*root).key > key {
		erase(&(*root).l, key)
	} else {
		erase(&(*root).r, key)
	}
}

// unite - за O(M log (N/M))
// Реализация Unite (X, Y).
// Пусть, не теряя общности, T1->Y > T2->Y, т.е. корень T1 будет корнем результата.
// Чтобы получить результат, нам нужно объединить деревья T1->L, T1->R и T2
// в два таких дерева, чтобы их можно было сделать сыновьями T1.
// Для этого вызовем Split (T2, T1->X), тем самым мы разобъём T2 на две половинки L и R,
// которые затем рекурсивно объединим с сыновьями T1: Union (T1->L, L) и Union (T1->R, R),
// тем самым мы построим левое и правое поддеревья результата.
func unite(l, r *treapNode) *treapNode {
	if l == nil || r == nil {
		if l == nil {
			return r
		}
		return l
	}

	if l.prior < r.prior {
		l, r = r, l
	}
	// r -> меньший приоритет
	// l -> больший приоритет
	// Мы возьмем l, а r мы поделим по ключу и дополним l данными из r,
	// а затем мы построим левое и правое поддеревья результата
	lroot := &treapNode{}
	rroot := &treapNode{}
	// Делим с меньшим приоритетом поддерево по ключу корня l (с наибольшим приоритетом)
	split(r, &lroot, &rroot, l.key)
	// Дополняем l.l данными из разделения
	l.l = unite(l.l, lroot)
	// Дополняем l.r данными из разделения
	l.r = unite(l.r, rroot)
	return l
}

func TestTreap(t *testing.T) {
	// a
	r := &treapNode{key: 5, prior: 5}
	b := &treapNode{key: 7, prior: 40}
	c := &treapNode{key: 8, prior: 3}
	d := &treapNode{key: 8, prior: 300}
	e := &treapNode{key: 1, prior: 301}

	x := unite(r, b)
	// b
	assert.Equal(t, x.prior, 40)
	assert.Equal(t, x.key, 7)
	// a
	assert.Equal(t, x.l.prior, 5)
	assert.Equal(t, x.l.key, 5)

	insert(&r, b)
	insert(&r, c)
	insert(&r, d)

	// d
	assert.Equal(t, r.prior, 300)
	assert.Equal(t, r.key, 8)
	// b
	assert.Equal(t, r.l.key, 7)
	assert.Equal(t, r.l.prior, 40)
	// c
	assert.Equal(t, r.l.r.key, 8)
	assert.Equal(t, r.l.r.prior, 3)
	// a
	assert.Equal(t, r.l.l.key, 5)
	assert.Equal(t, r.l.l.prior, 5)

	erase(&r, 7)
	// d
	assert.Equal(t, r.prior, 300)
	assert.Equal(t, r.key, 8)
	// a
	assert.Equal(t, r.l.key, 5)
	assert.Equal(t, r.l.prior, 5)
	// c
	assert.Equal(t, r.l.r.key, 8)
	assert.Equal(t, r.l.r.prior, 3)
	insert(&r, e)
	// e
	assert.Equal(t, r.prior, 301)
	assert.Equal(t, r.key, 1)

}
