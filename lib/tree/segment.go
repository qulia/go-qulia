package tree

import (
	"github.com/qulia/go-qulia/lib"
)

type SegmentTreeInterface interface {
	UpdateRange(start, end int, updateFunc lib.UpdateFunc)
	QueryRange(start, end int) interface{}
}

type SegmentTree struct {
	qFunc lib.QueryEvalFunc
	dFunc lib.DisjointValFunc
	root  *segmentTreeNode
}

func NewSegmentTree(qFunc lib.QueryEvalFunc, dFunc lib.DisjointValFunc) SegmentTreeInterface {
	st := SegmentTree{}
	st.qFunc = qFunc
	st.dFunc = dFunc
	st.root = &segmentTreeNode{
		r: rng{0, 1e9},
	}
	return &st
}

func (st SegmentTree) UpdateRange(start, end int, updateFunc lib.UpdateFunc) {
	st.root.updateRange(rng{start, end}, updateFunc, st.qFunc)
}

func (st SegmentTree) QueryRange(start, end int) interface{} {
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

type lazy struct {
	uFunc lib.UpdateFunc
	count int
}

type segmentTreeNode struct {
	r           rng
	data        interface{}
	lz          lazy
	left, right *segmentTreeNode
}

func (stn segmentTreeNode) isLeaf() bool {
	return stn.left == nil
}

func (stn *segmentTreeNode) lazyUpdate() {
	if stn.r.start != stn.r.end {
		if stn.isLeaf() {
			m := stn.r.mid()
			stn.left = &segmentTreeNode{r: rng{stn.r.start, m}, lz: lazy{}}
			stn.right = &segmentTreeNode{r: rng{m + 1, stn.r.end}, lz: lazy{}}
		}
	}

	if stn.lz.count != 0 {
		for i := 0; i < stn.lz.count; i++ {
			stn.data = stn.lz.uFunc(stn.data)
		}
		if !stn.isLeaf() {
			// propagate lazy
			stn.left.lz = stn.lz
			stn.right.lz = stn.lz
		}
		stn.lz.count = 0
	}
}

func (stn *segmentTreeNode) updateRange(r rng, uFunc lib.UpdateFunc, qFunc lib.QueryEvalFunc) {
	stn.lazyUpdate()
	if stn.r.disjoint(r) {
		return
	}

	if r.covers(stn.r) {
		stn.data = uFunc(stn.data)
		if !stn.isLeaf() {
			stn.left.lz.uFunc = uFunc
			stn.left.lz.count++
			stn.right.lz.uFunc = uFunc
			stn.right.lz.count++
		}
		return
	}

	stn.left.updateRange(r, uFunc, qFunc)
	stn.right.updateRange(r, uFunc, qFunc)
	stn.data = qFunc(stn.left.data, stn.right.data)
}

func (stn segmentTreeNode) queryRange(r rng, qFunc lib.QueryEvalFunc, dFunc lib.DisjointValFunc) interface{} {
	stn.lazyUpdate()
	if stn.r.disjoint(r) {
		return dFunc()
	}

	if stn.isLeaf() || r.covers(stn.r) {
		return stn.data
	}

	return qFunc(stn.left.queryRange(r, qFunc, dFunc), stn.right.queryRange(r, qFunc, dFunc))
}
