package unionfind

import (
	"testing"

	assert "github.com/stretchr/testify/require"
)

func TestUnionFind(t *testing.T) {
	unifTest := New([]int{1, 2, 3, 4, 5, 6, 7})
	unifTest.Union(1, 2)
	unifTest.Union(3, 4)
	unifTest.Union(4, 5)
	unifTest.Union(5, 6)
	unifTest.Union(1, 7)

	assert.True(t, unifTest.Find(1) == unifTest.Find(2))
	assert.True(t, unifTest.Find(1) == unifTest.Find(7))
	assert.True(t, unifTest.Find(3) == unifTest.Find(6))

	assert.Equal(t, 3, unifTest.Size(2)) // 1,2,7
	assert.Equal(t, 4, unifTest.Size(6)) //3,4,5,6
}
