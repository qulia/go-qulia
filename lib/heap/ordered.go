package heap

import (
	"github.com/qulia/go-qulia/lib"
	"golang.org/x/exp/constraints"
)

type ordered[T constraints.Ordered] struct {
	cch *customComp[lib.DefaultLesser[T]]
}

func (o *ordered[T]) Insert(elem T) {
	o.cch.Insert(lib.DefaultLesser[T]{Val: elem})
}

func (o *ordered[T]) Extract() T {
	return o.cch.Extract().Val
}

func (h ordered[T]) IsEmpty() bool {
	return h.cch.IsEmpty()
}

func (h ordered[T]) Size() int {
	return h.cch.Size()
}

func newOrdered[T constraints.Ordered](input []T, maxOnTop bool) Heap[T] {
	o := &ordered[T]{cch: newCustomComp[lib.DefaultLesser[T]](nil, maxOnTop)}
	for _, i := range input {
		o.Insert(i)
	}

	return o
}
