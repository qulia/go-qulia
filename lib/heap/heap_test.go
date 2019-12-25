package heap_test

import (
	"github.com/qulia/go-qulia/lib/heap"
	"github.com/stretchr/testify/assert"
	"testing"
)

var (
	intCompFunc = func(first, second interface{}) int { return first.(int) - second.(int) }
)
func TestMaxHeapBasicTest(t *testing.T) {
	input := []interface{}{3, 4, 7, 1, 0, -1}
	testHeap := heap.NewMaxHeap(input, intCompFunc)

	assert.Equal(t, 7, testHeap.Extract().(int))
	assert.Equal(t, 4, testHeap.Extract().(int))
	assert.Equal(t, 3, testHeap.Extract().(int))
	assert.Equal(t, 1, testHeap.Extract().(int))
	assert.Equal(t, 0, testHeap.Extract().(int))
	assert.Equal(t, -1, testHeap.Extract().(int))
	assert.Equal(t, nil, testHeap.Extract())
}

func TestMinHeapBasicTest(t *testing.T) {
	input := []interface{}{3, 4, 7, 1, 0, -1}
	testHeap := heap.NewMinHeap(input, intCompFunc)

	assert.Equal(t, -1, testHeap.Extract().(int))
	assert.Equal(t, 0, testHeap.Extract().(int))
	assert.Equal(t, 1, testHeap.Extract().(int))
	assert.Equal(t, 3, testHeap.Extract().(int))
	assert.Equal(t, 4, testHeap.Extract().(int))
	assert.Equal(t, 7, testHeap.Extract().(int))
	assert.Equal(t, nil, testHeap.Extract())
}

func TestHeapPushTest(t *testing.T) {
	testHeap := heap.NewMaxHeap(nil, intCompFunc)

	input := []interface{}{3, 4, 7, 1, 0, -1}
	for _,val := range input {
		testHeap.Insert(val)
	}

	assert.Equal(t, 7, testHeap.Extract().(int))
	assert.Equal(t, 4, testHeap.Extract().(int))
	assert.Equal(t, 3, testHeap.Extract().(int))
	assert.Equal(t, 1, testHeap.Extract().(int))
	assert.Equal(t, 0, testHeap.Extract().(int))
	assert.Equal(t, -1, testHeap.Extract().(int))
	assert.Equal(t, nil, testHeap.Extract())
}
