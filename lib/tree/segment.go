package tree

import "golang.org/x/exp/constraints"

type SegmentTree[T constraints.Ordered] interface {
	UpdateRange(start, end int, updateFunc UpdateFunc[T])
	QueryRange(start, end int) T
}

func NewSegmentTree[T constraints.Ordered](aggFunc AggregateFunc[T], identityEl T) *segmentTreeImpl[T] {
	return newSegmentTreeImpl(aggFunc, identityEl)
}

type AggregateFunc[T constraints.Ordered] func(a, b T) T
type UpdateFunc[T constraints.Ordered] func(current T) T
