package tree_test

import (
	"fmt"
	"testing"

	"github.com/qulia/go-qulia/lib/tree"
	"github.com/stretchr/testify/assert"
)

func TestTreeBasic(t *testing.T) {
	root := tree.NewNode(5)
	root.Left = tree.NewNode(3)
	root.Right = tree.NewNode(6)

	root.Left.Left = tree.NewNode(4)
	root.Right.Right = tree.NewNode(8)

	var result []string
	tree.Visit(root, func(elem interface{}) {
		var elemString string
		if elem == nil {
			elemString = "nil"
		} else {
			elemString = fmt.Sprintf("%d", elem.(int))
		}

		result = append(result, elemString)
	})

	assert.Equal(t, []string{"nil", "4", "nil", "3", "nil", "5", "nil", "6", "nil", "8", "nil"}, result)
}
