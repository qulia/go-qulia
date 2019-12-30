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
	evenNums.Add(2)
	evenNums.Add(4)
	evenNums.Add(6)
	evenNums.Add(6) // add twice, no-op
	evenNums.Add(8)

	oddNums := set.NewSet(lib.IntKeyFunc)
	oddNums.Add(1)
	oddNums.Add(3)
	oddNums.Add(3) // add twice, no-op
	oddNums.Add(5)
	oddNums.Add(7)
	oddNums.Add(9)

	primeNums := set.NewSet(lib.IntKeyFunc)
	primeNums.Add(2)
	primeNums.Add(3)
	primeNums.Add(5)
	primeNums.Add(7)

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

func sortSlice(input []interface{}) []interface{} {
	sort.Slice(input, func(i, j int) bool { return input[i].(int) < input[j].(int) })
	return input
}
