package tree_test

import (
	"sort"
	"testing"

	"github.com/qulia/go-qulia/lib"
	"github.com/qulia/go-qulia/lib/tree"
	"github.com/stretchr/testify/assert"
)

func TestBSTBasic(t *testing.T) {
	nums := []interface{}{8, 3, 10, 1, 6, 4, 7, 14, 13}
	bst := tree.NewBST(lib.IntCompFunc)
	for _, val := range nums {
		bst.Insert(val)
	}

	checkSorted(t, nums, bst)
	assert.Equal(t, 8, bst.Search(8).Data.(int))
	assert.Equal(t, 1, bst.Search(1).Data.(int))
	assert.Equal(t, 14, bst.Search(14).Data.(int))
	assert.Nil(t, bst.Search(16))
}

func checkSorted(t *testing.T, originalList []interface{}, bst tree.BSTInterface) {
	fromBST := bst.ToSlice()
	sort.Slice(originalList, func(i, j int) bool {
		return originalList[i].(int) < originalList[j].(int)
	})
	assert.Equal(t, originalList, fromBST)
}

func TestBSTEmpty(t *testing.T) {
	bst := tree.NewBST(lib.IntCompFunc)
	res := bst.ToSlice()
	var exp []interface{}
	assert.Equal(t, exp, res)

	assert.Nil(t, bst.Search(2))
	assert.Nil(t, bst.Search(nil))
}
