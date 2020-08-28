package tree

import (
	"math"
	"testing"

	"github.com/qulia/go-qulia/lib"
	assert "github.com/stretchr/testify/require"
)

func TestSegmentTreeMinEval(t *testing.T) {
	input := []int{1, 9, 5, 3, 7, 2, 4, 6, 8}
	st := NewSegmentTree(lib.IntQueryEvalMinFunc, func() interface{} { return math.MaxInt32 })
	initSt(t, input, st)

	assert.Equal(t, 3, st.QueryRange(2, 4).(int))
	assert.Equal(t, 1, st.QueryRange(0, 8).(int))
	assert.Equal(t, 1, st.QueryRange(0, 15).(int))

	// Increment [0,3] values by 3
	st.UpdateRange(0, 3, func(current interface{}) interface{} {
		if current != nil {
			return current.(int) + 3
		} else {
			return math.MaxInt32
		}
	})

	assert.Equal(t, 6, st.QueryRange(2, 4).(int))
	assert.Equal(t, 2, st.QueryRange(0, 8).(int))
	assert.Equal(t, 2, st.QueryRange(0, 15).(int))

}

func TestSegmentTreeSumEval(t *testing.T) {
	input := []int{1, 9, 5, 3, 7, 2, 4, 6, 8}
	st := NewSegmentTree(lib.IntQueryEvalSumFunc, func() interface{} { return 0 })
	initSt(t, input, st)

	assert.Equal(t, 15, st.QueryRange(2, 4).(int))
	assert.Equal(t, 45, st.QueryRange(0, 8).(int))
	assert.Equal(t, 45, st.QueryRange(0, 100).(int))
}

func initSt(t *testing.T, input []int, st SegmentTreeInterface) {
	for idx, val := range input {
		for count := 0; count < val; count++ {
			st.UpdateRange(idx, idx, func(current interface{}) interface{} {
				if current == nil {
					return 1
				} else {
					return current.(int) + 1
				}
			})
		}
		assert.Equal(t, val, st.QueryRange(idx, idx).(int))
	}
}
