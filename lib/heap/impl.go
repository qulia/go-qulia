package heap

import (
	"github.com/qulia/go-qulia/lib"
	"golang.org/x/exp/constraints"
)

type impl[T constraints.Ordered] struct {
	*fleximpl[lib.DefaultComparer[T]]
}

func newImpl[T constraints.Ordered](input []T, maxOnTop bool) Heap[T] {
	o := &impl[T]{newFlexImpl[lib.DefaultComparer[T]](nil, maxOnTop)}
	for _, i := range input {
		o.Insert(i)
	}

	return o
}

func (o *impl[T]) Insert(elem T) {
	o.fleximpl.Insert(lib.DefaultComparer[T]{Val: elem})
}

func (o *impl[T]) Extract() T {
	return o.fleximpl.Extract().Val
}
