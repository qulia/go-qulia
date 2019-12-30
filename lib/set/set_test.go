package set_test

import (
	"sort"
	"testing"

	"github.com/qulia/go-qulia/lib"

	"github.com/qulia/go-qulia/lib/set"
	"github.com/stretchr/testify/assert"
)

func TestSetBasic(t *testing.T) {
	evenNums := set.NewSet(lib.IntKeyFunc)
	addNums(evenNums, []int{2, 4, 6, 8})

	oddNums := set.NewSet(lib.IntKeyFunc)
	addNums(oddNums, []int{1, 3, 5, 7, 9})

	primeNums := set.NewSet(lib.IntKeyFunc)
	addNums(primeNums, []int{2, 3, 5, 7})

	nums := set.NewSet(lib.IntKeyFunc)
	nums.FromSlice([]interface{}{1, 2, 3, 4, 5, 6, 7, 8, 9})

	primeAndEvenNums := sortSlice(evenNums.Intersection(primeNums).ToSlice())
	oddAndPrimeEvenNums := sortSlice(primeNums.Intersection(oddNums).ToSlice())
	evenOrOddNums := sortSlice(evenNums.Union(oddNums).ToSlice())

	assert.Equal(t, evenNums.Size(), 4)
	assert.Equal(t, oddNums.Size(), 5)
	assert.Equal(t, 0, evenNums.Intersection(oddNums).Size())
	assert.Equal(t, []interface{}{2}, primeAndEvenNums)
	assert.Equal(t, []interface{}{3, 5, 7}, oddAndPrimeEvenNums)
	assert.Equal(t, []interface{}{1, 2, 3, 4, 5, 6, 7, 8, 9}, evenOrOddNums)

	assert.True(t, evenNums.IsSubsetOf(nums))
	assert.True(t, nums.IsSupersetOf(oddNums))

	nums.Remove(2)
	assert.False(t, evenNums.IsSubsetOf(nums))
	nums.Remove(3)
	assert.False(t, nums.IsSupersetOf(oddNums))
}

func TestSetDuplicate(t *testing.T) {
	nums := set.NewSet(lib.IntKeyFunc)
	addNums(nums, []int{8, 2, 4, 4, 6, 8})
	assert.Equal(t, 4, nums.Size())
}

func addNums(s *set.Set, nums []int) {
	for _, num := range nums {
		s.Add(num)
	}
}

func sortSlice(input []interface{}) []interface{} {
	sort.Slice(input, func(i, j int) bool { return input[i].(int) < input[j].(int) })
	return input
}
