package set

import "github.com/qulia/go-qulia/lib"

type SliceSet struct {
	buf []interface{}
	km  map[string]int
	kf  lib.KeyFunc
}

func NewSliceSet(kf lib.KeyFunc) *SliceSet {
	ss := SliceSet{kf: kf}
	ss.km = make(map[string]int)
	return &ss
}

func (ss SliceSet) GetSlice() []interface{} {
	return ss.buf
}

func (ss *SliceSet) Add(it interface{}) {
	k := ss.kf(it)
	ss.Remove(it)
	ss.buf = append(ss.buf, it)
	ss.km[k] = len(ss.buf) - 1
}

func (ss *SliceSet) Remove(it interface{}) {
	if !ss.ContainsKeyFor(it) {
		return
	}
	k := ss.kf(it)
	idx := ss.km[k]
	lastIdx := len(ss.buf) - 1
	// swap
	ss.buf[idx], ss.buf[lastIdx] = ss.buf[lastIdx], ss.buf[idx]
	ss.km[ss.kf(ss.buf[idx])] = idx
	ss.buf = ss.buf[:lastIdx]
	delete(ss.km, k)
}

func (ss SliceSet) ContainsKeyFor(it interface{}) bool {
	k := ss.kf(it)
	if _, ok := ss.km[k]; ok {
		return true
	}

	return false
}

func (ss SliceSet) GetItemForKey(it interface{}) interface{} {
	if !ss.ContainsKeyFor(it) {
		return nil
	}
	return ss.buf[ss.km[ss.kf(it)]]
}
