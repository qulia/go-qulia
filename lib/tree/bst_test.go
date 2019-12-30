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
}
