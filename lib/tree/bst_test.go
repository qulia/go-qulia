package tree_test

import (
	"sort"
	"testing"

	"github.com/qulia/go-qulia/v2/lib/tree"
	"github.com/stretchr/testify/assert"
)

func TestBSTBasic(t *testing.T) {
	nums := []int{8, 3, 10, 1, 6, 4, 7, 14, 13}
	bst := tree.NewBST[int]()
	for _, val := range nums {
		bst.Insert(val)
	}

	checkSorted(t, nums, bst)
	assert.Equal(t, 8, bst.Search(8).Data)
	assert.Equal(t, 1, bst.Search(1).Data)
	assert.Equal(t, 14, bst.Search(14).Data)
	assert.Nil(t, bst.Search(16))
}

func checkSorted(t *testing.T, originalList []int, bst tree.BST[int]) {
	fromBST := bst.ToSlice()
	sort.Slice(originalList, func(i, j int) bool {
		return originalList[i] < originalList[j]
	})
	assert.Equal(t, originalList, fromBST)
}

func TestBSTEmpty(t *testing.T) {
	bst := tree.NewBST[int]()
	res := bst.ToSlice()
	var exp []int
	assert.Equal(t, exp, res)

	assert.Nil(t, bst.Search(2))
}

func TestBST_Floor(t *testing.T) {
	nums := []int{0, 1, 5, 9, 1, 2, 3}
	bst := tree.NewBST[int]()
	for _, val := range nums {
		bst.Insert(val)
	}

	assert.Equal(t, 1, bst.Floor(1).Data)
	assert.Equal(t, 3, bst.Floor(3).Data)
	assert.Equal(t, 5, bst.Floor(6).Data)
	assert.Nil(t, bst.Floor(-1))
}

func TestBST_Ceiling(t *testing.T) {
	nums := []int{0, 1, 5, 9, 1, 2, 3}
	bst := tree.NewBST[int]()
	for _, val := range nums {
		bst.Insert(val)
	}

	assert.Equal(t, 1, bst.Ceiling(1).Data)
	assert.Equal(t, 3, bst.Ceiling(3).Data)
	assert.Equal(t, 5, bst.Ceiling(4).Data)
	assert.Equal(t, 9, bst.Ceiling(6).Data)
	assert.Nil(t, bst.Ceiling(10))
}
