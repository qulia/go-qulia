package set

import (
	"testing"

	"github.com/qulia/go-qulia/lib"
	assert "github.com/stretchr/testify/require"
)

func TestSliceSet(t *testing.T) {
	ss := NewSliceSet(lib.IntKeyFunc)
	for i := 0; i < 10; i++ {
		ss.Add(i)
	}

	assert.True(t, ss.ContainsKeyFor(4))
	assert.True(t, ss.ContainsKeyFor(7))
	assert.True(t, ss.ContainsKeyFor(3))
	assert.True(t, ss.ContainsKeyFor(8))

	ss.Remove(4)
	ss.Remove(7)
	ss.Remove(3)
	ss.Remove(8)

	assert.False(t, ss.ContainsKeyFor(4))
	assert.False(t, ss.ContainsKeyFor(7))
	assert.False(t, ss.ContainsKeyFor(3))
	assert.False(t, ss.ContainsKeyFor(8))

	assert.Nil(t, ss.GetItemForKey(4))
	assert.Equal(t, 5, ss.GetItemForKey(5))
}
