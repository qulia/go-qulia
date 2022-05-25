package tree

import (
	"math"
	"testing"

	assert "github.com/stretchr/testify/require"
	"golang.org/x/exp/constraints"
)

func TestSegmentTreeMinEval(t *testing.T) {
	input := []int{1, 9, 5, 3, 7, 2, 4, 6, 8}
	st := NewSegmentTree(MaxEvalFunc[int], math.MinInt32)
	initSt(t, input, st)

	assert.Equal(t, 7, st.QueryRange(2, 4))
	assert.Equal(t, 9, st.QueryRange(0, 8))
	assert.Equal(t, 9, st.QueryRange(0, 15))

	// Increment [0,3] values by 3
	st.UpdateRange(0, 3, func(current int) int {
		return current + 3
	})

	// 4, 12, 8, 6, 7, 2, 4, 6, 8
	assert.Equal(t, 8, st.QueryRange(2, 4))
	assert.Equal(t, 12, st.QueryRange(0, 8))
	assert.Equal(t, 12, st.QueryRange(0, 15))

	// dec [5, 8] values by 1
	st.UpdateRange(5, 8, func(current int) int {
		return current - 1
	})

	// 4, 12, 8, 6, 7, 1, 3, 5, 7
	assert.Equal(t, 8, st.QueryRange(2, 4))
	assert.Equal(t, 12, st.QueryRange(0, 8))
	assert.Equal(t, 12, st.QueryRange(0, 15))

}

func TestSegmentTreeSumEval(t *testing.T) {
	input := []int{1, 9, 5, 3, 7, 2, 4, 6, 8}
	st := NewSegmentTree(SumFunc[int], 0)
	initSt(t, input, st)

	assert.Equal(t, 15, st.QueryRange(2, 4))
	assert.Equal(t, 45, st.QueryRange(0, 8))
	assert.Equal(t, 45, st.QueryRange(0, 100))

	// Increment [6, 15] values by 5
	st.UpdateRange(6, 15, func(current int) int {
		return current + 5
	})

	// 1, 9, 5, 3, 7, 2, 9, 11, 13, 5, 5, 5, 5, 5, 5, 5
	assert.Equal(t, 95, st.QueryRange(0, 100))
}

func initSt(t *testing.T, input []int, st SegmentTree[int]) {
	for idx, val := range input {
		for count := 0; count < val; count++ {
			st.UpdateRange(idx, idx, func(current int) int {
				return current + 1
			})
		}
		assert.Equal(t, val, st.QueryRange(idx, idx))
	}
}

func MaxEvalFunc[T int](a, b T) T {
	if a > b {
		return a
	}
	return b
}

func SumFunc[T constraints.Ordered](a, b T) T {
	return a + b
}
