package queue_test

import (
	"testing"

	"github.com/qulia/go-qulia/v2/lib/queue"
	"github.com/stretchr/testify/assert"
)

func TestQueueBasic(t *testing.T) {
	testQueue := queue.NewQueue[int]()
	verify(t, testQueue)
}

func verify(t *testing.T, queue queue.Queue[int]) {
	queue.Enqueue(3)
	queue.Enqueue(7)
	queue.Enqueue(0)

	assert.Equal(t, 3, queue.Dequeue())
	assert.Equal(t, 7, queue.Dequeue())
	assert.Equal(t, 0, queue.Peek())
	assert.Equal(t, 0, queue.Dequeue())
	assert.True(t, queue.IsEmpty())
	assert.Panics(t, func() { queue.Dequeue() })
}
