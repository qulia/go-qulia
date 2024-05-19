package caching

import "github.com/qulia/go-qulia/v2/lib/common"

// Generic Cache interface
// values need to be Keyable
type Cache[T common.Keyable[K], K comparable] interface {
	// adds or updates the cache value
	Put(T)

	// Returns the [value,true] if key exists, [*new(T), false] if not
	Get(K) (T, bool)
}

// Least Recently Used (LRU) cache with capacity
// Operations are thread safe
// When capacity is reached oldest used item is evicted
func NewLRUCache[T common.Keyable[K], K comparable](capacity int) Cache[T, K] {
	return newLRUCache[T, K](capacity)
}
