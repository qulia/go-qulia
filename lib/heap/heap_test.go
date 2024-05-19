package heap_test

import (
	"fmt"
	"sort"
	"testing"

	"github.com/qulia/go-qulia/v2/lib/heap"
	"github.com/stretchr/testify/assert"
)

func TestMaxHeapBasic(t *testing.T) {
	input := []int{3, 4, 7, 1, 0, -1}
	testHeap := heap.NewMaxHeap(input)

	assert.Equal(t, 7, testHeap.Extract())
	assert.Equal(t, 4, testHeap.Extract())
	assert.Equal(t, 3, testHeap.Extract())
	assert.Equal(t, 1, testHeap.Extract())
	assert.Equal(t, 0, testHeap.Extract())
	assert.Equal(t, -1, testHeap.Extract())
	assert.Panics(t, func() { testHeap.Extract() })
}

func TestMinHeapBasic(t *testing.T) {
	input := []int{3, 4, 7, 1, 0, -1}
	testHeap := heap.NewMinHeap(input)

	assert.Equal(t, -1, testHeap.Extract())
	assert.Equal(t, 0, testHeap.Extract())
	assert.Equal(t, 1, testHeap.Extract())
	assert.Equal(t, 3, testHeap.Extract())
	assert.Equal(t, 4, testHeap.Extract())
	assert.Equal(t, 7, testHeap.Extract())
	assert.True(t, testHeap.IsEmpty())
	assert.Panics(t, func() { testHeap.Extract() })
}

func TestHeapPush(t *testing.T) {
	testHeap := heap.NewMaxHeap[int](nil)

	input := []int{3, 4, -1, 7, 7, 1, 0, -1, 4}
	for _, val := range input {
		testHeap.Insert(val)
	}

	assert.Equal(t, 7, testHeap.Extract())
	assert.Equal(t, 7, testHeap.Extract())
	assert.Equal(t, 4, testHeap.Extract())
	assert.Equal(t, 4, testHeap.Extract())
	assert.Equal(t, 3, testHeap.Extract())
	assert.Equal(t, 1, testHeap.Extract())
	assert.Equal(t, 0, testHeap.Extract())
	assert.Equal(t, -1, testHeap.Extract())
	assert.Equal(t, -1, testHeap.Extract())
	assert.True(t, testHeap.IsEmpty())
	assert.Panics(t, func() { testHeap.Extract() })
}

func TestMaxHeapWithRepeatingNumbers(t *testing.T) {
	input := []int{3, 4, -1, 7, 7, 1, 0, -1, 4}
	testHeap := heap.NewMaxHeap(input)

	assert.Equal(t, 7, testHeap.Extract())
	assert.Equal(t, 7, testHeap.Extract())
	assert.Equal(t, 4, testHeap.Extract())
	assert.Equal(t, 4, testHeap.Peek())
	assert.Equal(t, 4, testHeap.Extract())
	assert.Equal(t, 3, testHeap.Extract())
	assert.Equal(t, 1, testHeap.Extract())
	assert.Equal(t, 0, testHeap.Extract())
	assert.Equal(t, -1, testHeap.Extract())
	assert.Equal(t, -1, testHeap.Peek())
	assert.Equal(t, -1, testHeap.Extract())
	assert.True(t, testHeap.IsEmpty())
	assert.Panics(t, func() { testHeap.Extract() })
}

func TestHeapPushAllEqual(t *testing.T) {
	testHeap := heap.NewMaxHeap[int](nil)

	for i := 0; i < 10; i++ {
		testHeap.Insert(4)
	}

	for i := 0; i < 10; i++ {
		assert.Equal(t, 4, testHeap.Extract())
	}
	assert.True(t, testHeap.IsEmpty())
	assert.Panics(t, func() { testHeap.Extract() })
}

func TestMaxHeapAllEqual(t *testing.T) {
	testHeap := heap.NewMaxHeap([]int{4, 4, 4, 4, 4})
	for i := 0; i < 5; i++ {
		assert.Equal(t, 4, testHeap.Extract())
	}
	assert.True(t, testHeap.IsEmpty())
	assert.Panics(t, func() { testHeap.Extract() })
}

func TestMaxHeapSingleElem(t *testing.T) {
	testHeap := heap.NewMaxHeap([]int{4})
	assert.Equal(t, 4, testHeap.Extract())
	assert.True(t, testHeap.IsEmpty())
	assert.Panics(t, func() { testHeap.Extract() })
}

func TestHeapPushSingleElem(t *testing.T) {
	testHeap := heap.NewMaxHeap[int](nil)
	testHeap.Insert(4)
	assert.Equal(t, 4, testHeap.Extract())
	assert.True(t, testHeap.IsEmpty())
	assert.Panics(t, func() { testHeap.Extract() })
}

func TestHeapPushLargeInput(t *testing.T) {
	numsHeap := heap.NewMinHeap[int](nil)
	// Add to heap keeping the size fixed at k at most
	for i := 0; i < len(largeInput); i++ {
		numsHeap.Insert(largeInput[i])
	}

	var resultNums []int
	size := numsHeap.Size()
	for i := 0; i < size; i++ {
		resultNums = append(resultNums, numsHeap.Extract())
	}

	sortedInputBuffer := make([]int, len(largeInput))
	copy(sortedInputBuffer, largeInput)
	sort.Slice(sortedInputBuffer, func(i, j int) bool {
		return sortedInputBuffer[i] < sortedInputBuffer[j]
	})

	assert.Equal(t, sortedInputBuffer, resultNums)
}

func TestKthLargest(t *testing.T) {
	nums := largeInput
	numsHeap := heap.NewMinHeap[int](nil)
	k := 918
	// Add to heap keeping the size fixed at k at most
	for i := 0; i < len(nums); i++ {
		numsHeap.Insert(nums[i])
		if numsHeap.Size() > k {
			numsHeap.Extract()
		}
	}

	result := numsHeap.Extract()
	fmt.Printf("Result: %d\n", result)
	assert.Equal(t, 8221, result)
}
