package caching

import (
	"container/list"
	"sync"

	"github.com/qulia/go-qulia/lib/common"
)

type lruCache[T common.Keyable[K], K comparable] struct {
	cap   int
	m     map[K]*list.Element
	l     *list.List
	mutex sync.Mutex
}

type item[K comparable, T any] struct {
	key K
	val T
}

func newLRUCache[T common.Keyable[K], K comparable](capacity int) Cache[T, K] {
	return &lruCache[T, K]{
		cap:   capacity,
		m:     map[K]*list.Element{},
		l:     list.New(),
		mutex: sync.Mutex{},
	}
}

// Get implements Cache.
func (lc *lruCache[T, K]) Get(key K) (T, bool) {
	lc.mutex.Lock()
	defer lc.mutex.Unlock()
	if foundEl := lc.touchUnderLock(key); foundEl != nil {
		return foundEl.Value.(item[K, T]).val, true
	}

	return *new(T), false
}

// Put implements Cache.
func (lc *lruCache[T, K]) Put(t T) {
	lc.mutex.Lock()
	defer lc.mutex.Unlock()

	key := t.Key()
	if foundEl := lc.touchUnderLock(key); foundEl != nil {
		foundEl.Value = item[K, T]{key, t}
	} else {
		lc.cleanupUnderLock()
		lc.m[key] = lc.l.PushFront(item[K, T]{key, t})
	}
}

func (lc *lruCache[T, K]) cleanupUnderLock() {
	if lc.l.Len() == lc.cap {
		lasEl := lc.l.Back()
		key := lasEl.Value.(item[K, T]).key
		delete(lc.m, key)
		lc.l.Remove(lasEl)
	}
}

func (lc *lruCache[T, K]) touchUnderLock(key K) *list.Element {
	if foundEl, ok := lc.m[key]; ok {
		lc.l.MoveToFront(foundEl)
		return foundEl
	}

	return nil
}
