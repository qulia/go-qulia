package tree_test

import (
	"fmt"
	"testing"

	"github.com/qulia/go-qulia/v2/lib/tree"
	"github.com/stretchr/testify/assert"
)

func TestTreeBasic(t *testing.T) {
	root := tree.NewNode(5)
	root.Left = tree.NewNode(3)
	root.Right = tree.NewNode(6)

	root.Left.Left = tree.NewNode(4)
	root.Right.Right = tree.NewNode(8)

	var result []string
	tree.VisitInOrder(root, func(elem int) {
		result = append(result, fmt.Sprintf("%d", elem))
	})

	assert.Equal(t, []string{"4", "3", "5", "6", "8"}, result)
}
