package tree_test

import (
	"testing"

	"github.com/qulia/go-qulia/lib"
	"github.com/qulia/go-qulia/lib/tree"
	"github.com/stretchr/testify/assert"
)

func TestBSTBasic(t *testing.T) {
	nums := []int{8, 3, 10, 1, 6, 4, 7, 14, 13}
	bst := tree.NewBST(lib.IntCompFunc)
	for _, val := range nums {
		bst.Insert(val)
	}

	res := bst.ToSlice()
	assert.Equal(t, []interface{}{1, 3, 4, 6, 7, 8, 10, 13, 14}, res)

	assert.Equal(t, 8, bst.Search(8).Data.(int))
	assert.Equal(t, 1, bst.Search(1).Data.(int))
	assert.Equal(t, 14, bst.Search(14).Data.(int))
}

func TestBSTEmpty(t *testing.T) {
	bst := tree.NewBST(lib.IntCompFunc)
	res := bst.ToSlice()
	var exp []interface{}
	assert.Equal(t, exp, res)

	assert.Nil(t, bst.Search(2))
	assert.Nil(t, bst.Search(nil))
}
