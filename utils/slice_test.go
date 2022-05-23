package utils_test

import (
	"testing"

	"github.com/qulia/go-qulia/utils"

	"github.com/stretchr/testify/assert"
)

func TestSliceContains(t *testing.T) {
	elems := []int{3, 7, 1, 4, 4}
	elems2 := []int{3, 4}

	assert.True(t, utils.SliceContains(elems, elems2))
	assert.True(t, utils.SliceContains(elems, elems))
	assert.True(t, utils.SliceContains(elems2, elems2))
	assert.False(t, utils.SliceContains(elems2, elems))

	assert.False(t, utils.SliceContains([]int{}, elems2))
	assert.True(t, utils.SliceContains([]int{}, []int{}))
}

func TestSliceContainsElement(t *testing.T) {
	elems := []int{3, 7, 1, 4, 4}

	assert.True(t, utils.SliceContainsElement(elems, 3))
	assert.False(t, utils.SliceContainsElement(elems, 5))
	assert.False(t, utils.SliceContainsElement([]int{}, 3))
}
