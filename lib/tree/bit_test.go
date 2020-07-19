package tree

import (
	"testing"

	assert "github.com/stretchr/testify/require"
)

func TestBinaryIndexTree_Query(t *testing.T) {
	bit := NewBinaryIndexTree(10)
	for i := 1; i <= 10; i++ {
		bit.Update(i, i)
	}

	for i := 1; i <= 10; i++ {
		assert.Equal(t, i*(i+1)/2, bit.Query(i))
	}

	for i := 1; i <= 10; i++ {
		assert.Equal(t, i, bit.Query(i)-bit.Query(i-1))
	}
}
