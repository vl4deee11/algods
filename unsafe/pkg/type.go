package pkg

type ABC struct {
	i int8
	k int8
	v int8
}

func NewABC(i, k, v int8) *ABC {
	return &ABC{
		i: i,
		k: k,
		v: v,
	}
}
