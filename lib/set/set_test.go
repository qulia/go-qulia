package set_test

import (
	"sort"
	"testing"

	"github.com/qulia/go-qulia/v2/lib/set"
	"github.com/stretchr/testify/assert"
)

func TestSetBasic(t *testing.T) {
	evenNums := set.NewSet[int]()
	evenNums.FromSlice([]int{2, 4, 6, 8})

	oddNums := set.NewSet[int]()
	oddNums.FromSlice([]int{1, 3, 5, 7, 9})

	primeNums := set.NewSet[int]()
	primeNums.FromSlice([]int{2, 3, 5, 7})

	nums := set.NewSet[int]()
	nums.FromSlice(evenNums.Union(oddNums).ToSlice())

	primeAndEvenNums := evenNums.Intersection(primeNums).Keys()
	oddAndPrimeEvenNums := primeNums.Intersection(oddNums).Keys()
	evenOrOddNums := evenNums.Union(oddNums).Keys()

	sort.Ints(primeAndEvenNums)
	sort.Ints(oddAndPrimeEvenNums)
	sort.Ints(evenOrOddNums)

	assert.Equal(t, evenNums.Len(), 4)
	assert.Equal(t, oddNums.Len(), 5)
	assert.True(t, oddNums.Contains(3))
	assert.False(t, oddNums.Contains(2))
	assert.Equal(t, 0, evenNums.Intersection(oddNums).Len())
	assert.Equal(t, []int{2}, primeAndEvenNums)
	assert.Equal(t, []int{3, 5, 7}, oddAndPrimeEvenNums)
	assert.Equal(t, []int{1, 2, 3, 4, 5, 6, 7, 8, 9}, evenOrOddNums)

	assert.True(t, evenNums.IsSubsetOf(nums))
	assert.True(t, nums.IsSupersetOf(oddNums))

	nums.Remove(2)
	assert.False(t, evenNums.IsSubsetOf(nums))
	nums.Remove(3)
	assert.False(t, nums.IsSupersetOf(oddNums))
}

func TestSetDuplicate(t *testing.T) {
	nums := set.NewSet[int]()
	nums.FromSlice([]int{8, 2, 4, 4, 6, 8})
	assert.Equal(t, 4, nums.Len())
}

type employee struct {
	id            int
	name, surname string
	friends       []int
}

func (ek employee) Key() int {
	return ek.id
}

func TestCustomKey(t *testing.T) {
	emps := []employee{{1, "n1", "sn1", nil}, {2, "n2", "sn2", []int{1}}, {3, "n1", "sn1", nil}}
	s := set.NewSetFlex[employee, int]()
	for _, e := range emps {
		s.Add(e)
	}

	eids := s.Keys()
	sort.Ints(eids)
	assert.Equal(t, 3, s.Len())
	assert.Equal(t, []int{1, 2, 3}, eids)

	emps2 := []employee{{1, "n1", "sn1", nil}, {3, "n1", "sn1", nil}}
	s2 := set.NewSetFlex[employee, int]()
	s2.FromSlice(emps2)

	assert.Equal(t, s2.GetWithKey(1), emps[0])
	assert.Equal(t, s2.GetWithKey(5), *new(employee))
	assert.True(t, s2.IsSubsetOf(s))
	assert.False(t, s2.IsSupersetOf(s))
	assert.Equal(t, 3, s2.Union(s).Len())
	assert.Equal(t, 2, s2.Intersection(s).Len())

	s3 := set.NewSetFlex[employee, int]()
	assert.True(t, s3.FromSlice(s.ToSlice()).Union(s).IsSubsetOf(s))
}
