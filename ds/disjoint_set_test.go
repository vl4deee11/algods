package ds

import (
	"math/rand"
	"testing"
)

type Node struct {
	parent *Node
	rank   int
	Data   interface{}
}

func NewNode() *Node {
	s := &Node{}
	s.parent = s
	return s
}

func (e *Node) Find() *Node {
	for e.parent != e {
		e.parent = e.parent.parent
		e = e.parent
	}
	return e
}

func Union(e1, e2 *Node) {
	e1Root := e1.Find()
	e2Root := e2.Find()
	if e1Root == e2Root {
		return
	}

	switch {
	case e1Root.rank < e2Root.rank:
		e1Root.parent = e2Root
	case e1Root.rank > e2Root.rank:
		e2Root.parent = e1Root
	default:
		e2Root.parent = e1Root
		e1Root.rank++
	}
}

type Node_rec struct {
	parent *Node_rec
	rank   int
	data   int
}

func New(data int) *Node_rec {
	s := &Node_rec{}
	s.parent = s
	s.data = data
	return s
}

func Find_r(e *Node_rec) *Node_rec {
	if e.parent == e {
		return e
	}
	e.parent = Find_r(e.parent)
	return e.parent
}

func Union_r(e1, e2 *Node_rec) bool {
	r1 := Find_r(e1)
	r2 := Find_r(e2)
	if r1 == r2 {
		return false
	}

	switch {
	case r1.rank < r2.rank:
		r1.parent = r2
	case r1.rank > r2.rank:
		r2.parent = r1
	default:
		r2.parent = r1
		r1.rank++
	}
	return true
}

func TestEvenOdd(t *testing.T) {
	const N = 1000
	sets := make([]*Node, N)
	for i := 0; i < N; i++ {
		sets[i] = NewNode()
	}

	for i := 2; i < N; i += 2 {
		Union(sets[i], sets[i-2])
	}
	for i := 3; i < N; i += 2 {
		Union(sets[i], sets[i-2])
	}

	for i := 0; i < N*3; i++ {
		s1 := rand.Intn(N)
		s2 := rand.Intn(N)
		sameMod2 := s1%2 == s2%2
		sameRep := sets[s1].Find() == sets[s2].Find()
		if sameMod2 != sameRep {
			t.Fatalf("Should %d and %d lie in the same set?  The package incorrectly says %v.",
				s1, s2, sameRep)
		}
	}
}

func selectIndexes(n int) [][2]int {
	idxes := make([][2]int, n)
	if n < 2 {
		return idxes
	}
	for i := range idxes {
		idxes[i][0] = i
		if i == 0 {
			idxes[i][1] = rand.Intn(n)
		} else {
			idxes[i][1] = rand.Intn(i)
		}
	}
	return idxes
}

// go test -bench=. -gcflags '-l -N' -benchmem -cpu=1 -benchtime=1000000x
//BenchmarkUnion                   1000000               112.1 ns/op             0 B/op          0 allocs/op
//BenchmarkUnionFind               1000000               232.3 ns/op             0 B/op          0 allocs/op
// BenchmarkUnion measures the time to perform N union operations.
func BenchmarkUnion(b *testing.B) {
	b.StopTimer()
	elts := make([]*Node, b.N)
	for i := range elts {
		elts[i] = NewNode()
	}
	idxes := selectIndexes(b.N)
	b.StartTimer()
	for _, idx := range idxes {
		e1 := elts[idx[0]]
		e2 := elts[idx[1]]
		Union(e1, e2)
	}
}

// BenchmarkUnionFind measures the time to perform N union operations followed
// by N find operations.
func BenchmarkUnionFind(b *testing.B) {
	b.StopTimer()
	elts := make([]*Node, b.N)
	for i := range elts {
		elts[i] = NewNode()
	}
	idxes := selectIndexes(b.N)
	b.StartTimer()
	for _, idx := range idxes {
		e1 := elts[idx[0]]
		e2 := elts[idx[1]]
		Union(e1, e2)
	}
	for _, e := range elts {
		_ = e.Find()
	}
}
