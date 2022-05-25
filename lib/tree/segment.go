package tree

import "golang.org/x/exp/constraints"

type SegmentTree[T constraints.Ordered] interface {
	UpdateRange(start, end int, updateFunc UpdateFunc[T])
	QueryRange(start, end int) T
}

func NewSegmentTree[T constraints.Ordered](aggFunc AggregateFunc[T], identityEl T) *orderedSegmentTree[T] {
	st := orderedSegmentTree[T]{aggFunc, identityEl, &segmentTreeNode[T]{r: rng{0, 1e9}}}
	return &st
}

type AggregateFunc[T constraints.Ordered] func(a, b T) T
type UpdateFunc[T constraints.Ordered] func(current T) T
