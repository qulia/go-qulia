package set

import "github.com/qulia/go-qulia/lib"

type SliceSet[T lib.Keyable[K], K comparable] struct {
	buf []T  // holds items
	km  map[K]int  // key to buf index
}

func NewSliceSet[T lib.Keyable[K], K comparable] () *SliceSet[T, K] {
	return &SliceSet[T, K]{km:map[K]int{}}
}

func (ss SliceSet[T, K]) GetSlice() []T {
	return ss.buf
}

func (ss *SliceSet[T, K]) Add(it T) {
	ss.Remove(it)
	ss.buf = append(ss.buf, it)
	ss.km[it.Key()] = len(ss.buf) - 1
}

func (ss *SliceSet[T, K]) Remove(it T) {
	ss.RemoveKey(it.Key())
}

func (ss *SliceSet[T, K]) RemoveKey(key K) {
	if !ss.ContainsKey(key) {
		return
	}
	idx := ss.km[key]
	lastIdx := len(ss.buf) - 1
	// swap
	ss.buf[idx], ss.buf[lastIdx] = ss.buf[lastIdx], ss.buf[idx]
	ss.km[ss.buf[idx].Key()] = idx
	ss.buf = ss.buf[:lastIdx]
	delete(ss.km, key)
}

func (ss SliceSet[T, K]) Contains(it T) bool {
	return ss.ContainsKey(it.Key())
}

func (ss SliceSet[T, K]) ContainsKey(key K) bool {
	_, ok := ss.km[key]
	return ok
}

func (ss SliceSet[T, K]) GetItemWithKey(key K) (T,bool) {
	if !ss.ContainsKey(key) {
		return ss.buf[ss.km[key]], false
	}
	return ss.buf[ss.km[key]], true
}
