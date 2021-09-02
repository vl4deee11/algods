package ds

import "testing"

// -1 as spec var
type stack interface {
	push(int)
	top() int
	pop() int
}

// stackOnMap inst based on map
type stackOnMap struct {
	m map[int]int
	t int
}

func (s *stackOnMap) push(e int) {
	s.t++
	s.m[s.t] = e
}

func (s *stackOnMap) top() int {
	if s.t == -1 {
		return -1
	}
	return s.m[s.t]
}

func (s *stackOnMap) pop() int {
	r := s.top()
	if r == -1 {
		return r
	}
	delete(s.m, s.t)
	s.t--
	return r
}

// stackOnSlice of inst based on slice
type stackOnSlice struct {
	m []int
}

func (s *stackOnSlice) push(e int) {
	s.m = append(s.m, e)
}

func (s *stackOnSlice) top() int {
	if len(s.m) == 0 {
		return -1
	}
	return s.m[len(s.m)-1]
}

func (s *stackOnSlice) pop() int {
	r := s.top()
	if r == -1 {
		return r
	}
	s.m = s.m[:len(s.m)-1]
	return r
}

type n struct {
	e      int
	nx, pv *n
}

// stackOn2Llist of inst based on doubly linked list
type stackOn2Llist struct {
	m *n
}

func (s *stackOn2Llist) push(e int) {
	if s.m == nil {
		s.m = &n{e: e}
		return
	}

	s.m.nx = &n{e: e, pv: s.m}
	s.m = s.m.nx
}

func (s *stackOn2Llist) top() int {
	if s.m == nil {
		return -1
	}

	return s.m.e
}

func (s *stackOn2Llist) pop() int {
	r := s.top()
	if r == -1 {
		return r
	}
	s.m = s.m.pv
	return r
}

// go test -bench=. -gcflags '-l -N' -benchmem -cpu=1 -benchtime=1000000x
//Benchmark_StackOnMap_Push        1000000               209 ns/op              87 B/op          0 allocs/op
//Benchmark_StackOnSlice_Push      1000000                18.6 ns/op            45 B/op          0 allocs/op
//Benchmark_StackOn2LList_Push     1000000               106 ns/op              32 B/op          1 allocs/op
//Benchmark_StackOnMap_Pop         1000000               108 ns/op               0 B/op          0 allocs/op
//Benchmark_StackOnSlice_Pop       1000000                 8.31 ns/op            0 B/op          0 allocs/op
//Benchmark_StackOn2LList_Pop      1000000                 6.32 ns/op            0 B/op          0 allocs/op

func Benchmark_StackOnMap_Push(b *testing.B) {
	s := &stackOnMap{m: map[int]int{}, t: -1}
	for i := 0; i < b.N; i++ {
		s.push(i)
	}
}

func Benchmark_StackOnSlice_Push(b *testing.B) {
	s := new(stackOnSlice)
	for i := 0; i < b.N; i++ {
		s.push(i)
	}
}

func Benchmark_StackOn2LList_Push(b *testing.B) {
	s := new(stackOn2Llist)
	for i := 0; i < b.N; i++ {
		s.push(i)
	}
}

var Res int

func Benchmark_StackOnMap_Pop(b *testing.B) {
	b.StopTimer()
	s := &stackOnMap{m: map[int]int{}, t: -1}
	r2 := 0
	for i := 0; i < b.N; i++ {
		s.push(i)
	}
	s.push(0)
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		r2 = s.pop()
	}
	Res = r2
}

func Benchmark_StackOnSlice_Pop(b *testing.B) {
	b.StopTimer()
	s := new(stackOnSlice)
	r2 := 0
	for i := 0; i < b.N; i++ {
		s.push(i)
	}
	s.push(0)
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		r2 = s.pop()
	}
	Res = r2
}

func Benchmark_StackOn2LList_Pop(b *testing.B) {
	b.StopTimer()
	s := new(stackOn2Llist)
	r2 := 0
	for i := 0; i < b.N; i++ {
		s.push(i)
	}
	s.push(0)
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		r2 = s.pop()
	}
	Res = r2
}

func Test_StackOnMap(t *testing.T) {
	s := &stackOnMap{m: map[int]int{}, t: -1}
	testStack(s, t)
}

func Test_StackOnSlice(t *testing.T) {
	s := new(stackOnSlice)
	testStack(s, t)
}

func Test_StackOn2LList(t *testing.T) {
	s := new(stackOn2Llist)
	testStack(s, t)
}

func testStack(s stack, t *testing.T) {
	if s.pop() != -1 {
		t.Errorf("s.top() != -1")
	}

	if s.top() != -1 {
		t.Errorf("s.top() != -1")
	}

	s.push(1)
	s.push(2)
	s.push(456)
	if s.top() != 456 {
		t.Errorf("s.top() != 456")
	}

	s.pop()
	if s.top() != 2 {
		t.Errorf("s.top() != 2")
	}

	s.push(23)
	s.pop()
	s.pop()
	if s.top() != 1 {
		t.Errorf("s.top() != 1")
	}
}
