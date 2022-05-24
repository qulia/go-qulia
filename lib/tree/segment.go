package tree

import "golang.org/x/exp/constraints"

type SegmentTree[T constraints.Ordered] interface {
	UpdateRange(start, end int, updateFunc UpdateFunc[T])
	QueryRange(start, end int) T
}

type orderedSegmentTree[T constraints.Ordered] struct {
	qFunc QueryEvalFunc[T]   // used to aggregate result while navigating the tree
	dFunc DisjointValFunc[T] // used to get a value when the range is disjoint
	root  *segmentTreeNode[T]
}

func NewSegmentTree[T constraints.Ordered](qFunc QueryEvalFunc[T], dFunc DisjointValFunc[T]) *orderedSegmentTree[T] {
	st := orderedSegmentTree[T]{}
	st.qFunc = qFunc
	st.dFunc = dFunc
	st.root = &segmentTreeNode[T]{
		r: rng{0, 1e9},
	}
	return &st
}

func (st orderedSegmentTree[T]) UpdateRange(start, end int, updateFunc UpdateFunc[T]) {
	st.root.updateRange(rng{start, end}, updateFunc, st.qFunc)
}

func (st orderedSegmentTree[T]) QueryRange(start, end int) T {
	return st.root.queryRange(rng{start, end}, st.qFunc, st.dFunc)
}

type rng struct {
	start, end int
}

func (r rng) mid() int {
	return (r.start + r.end) / 2
}

func (r rng) disjoint(other rng) bool {
	return r.end < other.start || other.end < r.start
}

func (r rng) covers(other rng) bool {
	return r.start <= other.start && r.end >= other.end
}

type lazy[T constraints.Ordered] struct {
	uFunc UpdateFunc[T]
	count int
}

type segmentTreeNode[T constraints.Ordered] struct {
	r           rng
	data        T
	left, right *segmentTreeNode[T]
}

func (stn segmentTreeNode[T]) isLeaf() bool {
	return stn.left == nil
}

func (stn *segmentTreeNode[T]) split() {
	if stn.r.start != stn.r.end {
		if stn.isLeaf() {
			m := stn.r.mid()
			stn.left = &segmentTreeNode[T]{r: rng{stn.r.start, m}}
			stn.right = &segmentTreeNode[T]{r: rng{m + 1, stn.r.end}}
		}
	}
}

func (stn *segmentTreeNode[T]) updateRange(r rng, uFunc UpdateFunc[T], qFunc QueryEvalFunc[T]) {
	stn.split()
	if stn.r.disjoint(r) {
		return
	}

	if stn.isLeaf() {
		stn.data = uFunc(stn.data)
		return
	}

	stn.left.updateRange(r, uFunc, qFunc)
	stn.right.updateRange(r, uFunc, qFunc)
	stn.data = qFunc(stn.left.data, stn.right.data)
}

func (stn segmentTreeNode[T]) queryRange(r rng, qFunc QueryEvalFunc[T], dFunc DisjointValFunc[T]) T {
	stn.split()
	if stn.r.disjoint(r) {
		return dFunc()
	}

	if stn.isLeaf() || r.covers(stn.r) {
		return stn.data
	}

	return qFunc(stn.left.queryRange(r, qFunc, dFunc), stn.right.queryRange(r, qFunc, dFunc))
}

type QueryEvalFunc[T constraints.Ordered] func(a, b T) T
type UpdateFunc[T constraints.Ordered] func(current T) T
type DisjointValFunc[T constraints.Ordered] func() T
