package tree

import (
	"testing"

	assert "github.com/stretchr/testify/require"
)

func TestBinaryIndexTree_Query(t *testing.T) {
	bit := NewBinaryIndexTree(10)
	for i := 1; i <= 10; i++ {
		// 0:1, 1:2, 2:3, ...
		bit.Update(i-1, i)
	}

	for i := 1; i <= 10; i++ {
		assert.Equal(t, i*(i+1)/2, bit.Sum(i-1))
	}

	for i := 1; i <= 10; i++ {
		assert.Equal(t, i, bit.Sum(i-1)-bit.Sum(i-2))
	}

	bit.Update(5, 8)                     // update 6 to 8
	assert.Equal(t, 7*8/2+2, bit.Sum(6)) // (sum 1 to 7) + 2
}
