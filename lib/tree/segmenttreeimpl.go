package tree

import "golang.org/x/exp/constraints"

type segmentTreeImpl[T constraints.Ordered] struct {
	aggFunc   AggregateFunc[T] // aggregate func for range
	identiyEl T                // identity element for aggregate func. e.g. agg: sum, identiy: 0
	root      *segmentTreeNode[T]
}

func newSegmentTreeImpl[T constraints.Ordered](aggFunc AggregateFunc[T], identityEl T) *segmentTreeImpl[T] {
	st := segmentTreeImpl[T]{aggFunc, identityEl, &segmentTreeNode[T]{r: rng{0, 1e9}}}
	return &st
}

func (st segmentTreeImpl[T]) UpdateRange(start, end int, updateFunc UpdateFunc[T]) {
	st.root.updateRange(rng{start, end}, updateFunc, st.aggFunc)
}

func (st segmentTreeImpl[T]) QueryRange(start, end int) T {
	return st.root.queryRange(rng{start, end}, st.aggFunc, st.identiyEl)
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

func (stn *segmentTreeNode[T]) updateRange(r rng, uFunc UpdateFunc[T], qFunc AggregateFunc[T]) {
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

func (stn segmentTreeNode[T]) queryRange(r rng, aggFunc AggregateFunc[T], identityEl T) T {
	stn.split()
	if stn.r.disjoint(r) {
		return identityEl
	}

	if stn.isLeaf() || r.covers(stn.r) {
		return stn.data
	}

	return aggFunc(stn.left.queryRange(r, aggFunc, identityEl), stn.right.queryRange(r, aggFunc, identityEl))
}
