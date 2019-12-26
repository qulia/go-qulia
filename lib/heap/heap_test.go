package heap_test

import (
	"github.com/qulia/go-qulia/lib/heap"
	"github.com/stretchr/testify/assert"
	"testing"
)

var (
	intCompFunc = func(first, second interface{}) int {
		firstInt :=first.(int)
		secondInt := second.(int)
		if firstInt < secondInt {
			return -1
		} else if firstInt > secondInt {
			return 1
		} else {
			return 0
		}
	}
)

func TestMaxHeapBasic(t *testing.T) {
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

func TestMinHeapBasic(t *testing.T) {
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

func TestHeapPush(t *testing.T) {
	testHeap := heap.NewMaxHeap(nil, intCompFunc)

	input := []interface{}{3, 4, -1, 7, 7, 1, 0, -1, 4}
	for _,val := range input {
		testHeap.Insert(val)
	}

	assert.Equal(t, 7, testHeap.Extract().(int))
	assert.Equal(t, 7, testHeap.Extract().(int))
	assert.Equal(t, 4, testHeap.Extract().(int))
	assert.Equal(t, 4, testHeap.Extract().(int))
	assert.Equal(t, 3, testHeap.Extract().(int))
	assert.Equal(t, 1, testHeap.Extract().(int))
	assert.Equal(t, 0, testHeap.Extract().(int))
	assert.Equal(t, -1, testHeap.Extract().(int))
	assert.Equal(t, -1, testHeap.Extract().(int))
	assert.Equal(t, nil, testHeap.Extract())
}

func TestMaxHeapWithRepeatingNumbers(t *testing.T) {
	input := []interface{}{3, 4, -1, 7, 7, 1, 0, -1, 4}
	testHeap := heap.NewMaxHeap(input, intCompFunc)

	assert.Equal(t, 7, testHeap.Extract().(int))
	assert.Equal(t, 7, testHeap.Extract().(int))
	assert.Equal(t, 4, testHeap.Extract().(int))
	assert.Equal(t, 4, testHeap.Extract().(int))
	assert.Equal(t, 3, testHeap.Extract().(int))
	assert.Equal(t, 1, testHeap.Extract().(int))
	assert.Equal(t, 0, testHeap.Extract().(int))
	assert.Equal(t, -1, testHeap.Extract().(int))
	assert.Equal(t, -1, testHeap.Extract().(int))
	assert.Equal(t, nil, testHeap.Extract())
}

func TestHeapPushAllEqual(t *testing.T) {
	testHeap := heap.NewMaxHeap(nil, intCompFunc)

	for i := 0; i < 10; i++ {
		testHeap.Insert(4)
	}

	for i := 0; i < 10; i++ {
		assert.Equal(t, 4, testHeap.Extract().(int))
	}
	assert.Equal(t, nil, testHeap.Extract())
}

func TestMaxHeapAllEqual(t *testing.T) {
	testHeap := heap.NewMaxHeap([]interface{}{4, 4, 4, 4,4}, intCompFunc)
	for i := 0; i < 5; i++ {
		assert.Equal(t, 4, testHeap.Extract().(int))
	}
	assert.Equal(t, nil, testHeap.Extract())
}

func TestMaxHeapSingleElem(t *testing.T) {
	testHeap := heap.NewMaxHeap([]interface{}{4}, intCompFunc)
	assert.Equal(t, 4, testHeap.Extract().(int))
	assert.Equal(t, nil, testHeap.Extract())
}

func TestHeapPushSingleElem(t *testing.T) {
	testHeap := heap.NewMaxHeap(nil, intCompFunc)
	testHeap.Insert(4)
	assert.Equal(t, 4, testHeap.Extract().(int))
	assert.Equal(t, nil, testHeap.Extract())
}