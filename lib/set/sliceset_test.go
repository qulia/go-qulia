package set

import (
	"testing"

	"github.com/qulia/go-qulia/lib"
	assert "github.com/stretchr/testify/require"
)

func TestSliceSet(t *testing.T) {
	ss := NewSliceSet[lib.DefaultKeyable[int], int]()
	for i := 0; i < 10; i++ {
		ss.Add(lib.DefaultKeyable[int]{Val:i})
	}

	assert.True(t, ss.ContainsKey(4))
	assert.True(t, ss.ContainsKey(7))
	assert.True(t, ss.ContainsKey(3))
	assert.True(t, ss.ContainsKey(8))

	ss.RemoveKey(4)
	ss.RemoveKey(7)
	ss.RemoveKey(3)
	ss.RemoveKey(8)

	assert.False(t, ss.ContainsKey(4))
	assert.False(t, ss.ContainsKey(7))
	assert.False(t, ss.ContainsKey(3))
	assert.False(t, ss.ContainsKey(8))

	_, exists := ss.GetItemWithKey(4)
	assert.False(t, exists)
	_, exists = ss.GetItemWithKey(5)
	assert.True(t, exists)
}
