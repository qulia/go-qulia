package caching_test

import (
	"fmt"
	"testing"

	"github.com/qulia/go-qulia/caching"
	"github.com/qulia/go-qulia/lib/common"
	"github.com/stretchr/testify/assert"
)

type todo struct {
	id          int
	title       string
	description string
}

func (t todo) Key() int {
	return t.id
}

func TestLRUCacheBasic(t *testing.T) {
	cap := 10
	lc := caching.NewLRUCache[todo, int](10)

	todos := []todo{}
	for i := 0; i < cap*2; i++ {
		todos = append(todos, todo{
			id:          i,
			title:       fmt.Sprintf("Title for item %d", i),
			description: fmt.Sprintf("Descriptions for item %d", i),
		})
	}
	addIndex := cap
	for i := 0; i < cap; i++ {
		lc.Put(todos[i])
	}

	// 9 -> 8 -> 7 -> 6 -> 5 -> 4 -> 3 -> 2 -> 1 -> 0
	checkExists(todos, 0, addIndex, lc, t)

	// Last accessed one is the 0th at this point
	lc.Put(todos[addIndex]) // this should evict 0th
	addIndex++
	// 10 -> 9 -> 8 -> 7 -> 6 -> 5 -> 4 -> 3 -> 2 -> 1
	checkNotExists(0, lc, t)

	// Last accessed one is the 1st at this point
	// If we get it should not be evicted after adding another
	lc.Get(1)
	lc.Put(todos[addIndex])
	addIndex++
	checkNotExists(2, lc, t)
	// 1 -> 11 -> 10-> 9 -> 8 -> 7 -> 6 -> 5 -> 4 -> 3

	// Last accessed one is the 3rd at this point
	// If we do a put should not be evicted after adding another
	todos[3].title = "New title"
	lc.Put(todos[3])
	lc.Put(todos[addIndex])
	addIndex++
	checkNotExists(4, lc, t)
	// 12 -> 3 -> 1 -> 11 -> 10-> 9 -> 8 -> 7 -> 6 -> 5

	// Check the final state
	checkExists(todos, 1, 1, lc, t)
	checkExists(todos, 3, 3, lc, t)
	checkExists(todos, 5, addIndex, lc, t)
}

func checkNotExists[T common.Keyable[int]](key int, lc caching.Cache[T, int], t *testing.T) {
	val, exists := lc.Get(key)
	assert.Equal(t, *new(todo), val)
	assert.False(t, exists)
}

func checkExists[T common.Keyable[int]](
	todos []todo, start, end int, lc caching.Cache[T, int], t *testing.T,
) {
	for i := start; i < end; i++ {
		val, exists := lc.Get(i)
		assert.Equal(t, val, todos[i])
		assert.True(t, exists)
	}
}
