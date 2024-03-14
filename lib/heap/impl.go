package heap

import (
	"github.com/qulia/go-qulia/lib/common"
	"golang.org/x/exp/constraints"
)

type impl[T constraints.Ordered] struct {
	*fleximpl[common.DefaultComparer[T]]
}

func newImpl[T constraints.Ordered](input []T, maxOnTop bool) Heap[T] {
	o := &impl[T]{newFlexImpl[common.DefaultComparer[T]](nil, maxOnTop)}
	for _, i := range input {
		o.Insert(i)
	}

	return o
}

func (o *impl[T]) Insert(elem T) {
	o.fleximpl.Insert(common.DefaultComparer[T]{Val: elem})
}

func (o *impl[T]) Extract() T {
	return o.fleximpl.Extract().Val
}

func (o *impl[T]) Peek() T {
	return o.fleximpl.Peek().Val
}
