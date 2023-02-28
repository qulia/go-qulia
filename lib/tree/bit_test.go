package tree

import (
	"testing"

	assert "github.com/stretchr/testify/require"
)

func TestBinaryIndexTree_Query(t *testing.T) {
	bit := NewBinaryIndexTree(10)
	for i := 1; i <= 10; i++ {
		bit.Update(i-1, i)
	}

	for i := 1; i <= 10; i++ {
		assert.Equal(t, i*(i+1)/2, bit.Sum(i-1))
	}

	for i := 1; i <= 10; i++ {
		assert.Equal(t, i, bit.Sum(i-1)-bit.Sum(i-2))
	}
}
