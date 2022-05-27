package heap

import (
	"github.com/qulia/go-qulia/lib"
	"golang.org/x/exp/constraints"
)

type impl[T constraints.Ordered] struct {
	cch *fleximpl[lib.DefaultLesser[T]]
}

func newImpl[T constraints.Ordered](input []T, maxOnTop bool) Heap[T] {
	o := &impl[T]{cch: newFlexImpl[lib.DefaultLesser[T]](nil, maxOnTop)}
	for _, i := range input {
		o.Insert(i)
	}

	return o
}

func (o *impl[T]) Insert(elem T) {
	o.cch.Insert(lib.DefaultLesser[T]{Val: elem})
}

func (o *impl[T]) Extract() T {
	return o.cch.Extract().Val
}

func (h impl[T]) IsEmpty() bool {
	return h.cch.IsEmpty()
}

func (h impl[T]) Size() int {
	return h.cch.Size()
}
