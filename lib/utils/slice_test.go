package utils_test

import (
	"testing"

	"github.com/qulia/go-qulia/lib"
	"github.com/qulia/go-qulia/lib/utils"
	"github.com/stretchr/testify/assert"
)

func TestSliceContains(t *testing.T) {
	elems := []interface{}{3, 7, 1, 4, 4}
	elems2 := []interface{}{3, 4}

	assert.True(t, utils.SliceContains(elems, elems2, lib.IntKeyFunc))
	assert.True(t, utils.SliceContains(elems, elems, lib.IntKeyFunc))
	assert.True(t, utils.SliceContains(elems2, elems2, lib.IntKeyFunc))
	assert.False(t, utils.SliceContains(elems2, elems, lib.IntKeyFunc))

	assert.False(t, utils.SliceContains([]interface{}{}, elems2, lib.IntKeyFunc))
	assert.True(t, utils.SliceContains([]interface{}{}, []interface{}{}, lib.IntKeyFunc))
}

func TestSliceContainsElement(t *testing.T) {
	elems := []interface{}{3, 7, 1, 4, 4}

	assert.True(t, utils.SliceContainsElement(elems, 3, lib.IntKeyFunc))
	assert.False(t, utils.SliceContainsElement(elems, 5, lib.IntKeyFunc))
	assert.False(t, utils.SliceContainsElement([]interface{}{}, 3, lib.IntKeyFunc))
}
