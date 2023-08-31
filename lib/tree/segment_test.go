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

func TestSegmentTreeEventAvailability(t *testing.T) {
	events := [][]int{
		{1, 4},
		{6, 8},
		{2, 3},
		{5, 7},
	}
	people := map[string][][]int{
		"p1": {{1, 2}},
		"p2": {{3, 4}, {5, 8}},
		"p3": {{6, 9}},
		"p4": {{1, 9}},
	}

	// canAttend is true if the slot is available for the event for that attendee
	// does not take into account blocking, see the next case
	canAttendExpected := map[string]map[int]bool{
		"p1": {0: false, 1: false, 2: false, 3: false},
		"p2": {0: false, 1: true, 2: false, 3: true},
		"p3": {0: false, 1: true, 2: false, 3: false},
		"p4": {0: true, 1: true, 2: true, 3: true},
	}

	peopleAvailibility := map[string]SegmentTree[int]{}
	for p, av := range people {
		peopleAvailibility[p] = NewSegmentTree(SumFunc[int], 0)
		for _, a := range av {
			peopleAvailibility[p].UpdateRange(a[0], a[1], func(current int) int {
				return current + 1
			})
		}
	}

	for e := range events {
		for p := range people {
			duration := events[e][1] - events[e][0] + 1
			actual := peopleAvailibility[p].QueryRange(events[e][0], events[e][1]) == duration
			expected := canAttendExpected[p][e]
			t.Logf("Event %d, person %s: expected %t, actual %t", e, p, expected, actual)
			assert.Equal(t, expected, actual)
		}
	}
}

func TestSegmentTreeEventScheduling(t *testing.T) {
	events := [][]int{
		{1, 4},
		{6, 8},
		{2, 3},
		{5, 7},
	}
	people := map[string][][]int{
		"p1": {{1, 2}},
		"p2": {{3, 4}, {5, 8}},
		"p3": {{6, 9}},
		"p4": {{1, 9}},
	}

	// for this case events are processed in order and if the attendee can attend
	// it will block off that time; the next event will be evaluated based on updated
	// "calendar"
	canAttendExpected := map[string]map[int]bool{
		"p1": {0: false, 1: false, 2: false, 3: false},
		"p2": {0: false, 1: true, 2: false, 3: false /*blocked by [6, 8] event now*/},
		"p3": {0: false, 1: true, 2: false, 3: false},
		"p4": {0: true, 1: true, 2: false /* blocked by [1,4] event now*/, 3: false /*blocked by [6,8] event now*/},
	}

	peopleAvailibility := map[string]SegmentTree[int]{}
	for p, av := range people {
		peopleAvailibility[p] = NewSegmentTree(SumFunc[int], 0)
		for _, a := range av {
			peopleAvailibility[p].UpdateRange(a[0], a[1], func(current int) int {
				// available means 1, blocked means 0
				return current + 1
			})
		}
	}

	for e := range events {
		for p := range people {
			duration := events[e][1] - events[e][0] + 1
			actual := peopleAvailibility[p].QueryRange(events[e][0], events[e][1]) == duration
			expected := canAttendExpected[p][e]
			t.Logf("Event %d, person %s: expected %t, actual %t", e, p, expected, actual)
			assert.Equal(t, expected, actual)
			if expected {
				// block of the time for the attendee
				peopleAvailibility[p].UpdateRange(events[e][0], events[e][1], func(current int) int {
					return 0
				})
			}
		}
	}
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
