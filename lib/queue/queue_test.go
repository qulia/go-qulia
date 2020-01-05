package queue_test

import (
	"testing"

	"github.com/qulia/go-qulia/lib/queue"

	"github.com/stretchr/testify/assert"
)

func TestQueueBasic(t *testing.T) {
	testQueue := queue.Queue{}
	verify(t, &testQueue)
}

func verify(t *testing.T, queue queue.Interface) {
	queue.Enqueue(3)
	queue.Enqueue("strings")
	queue.Enqueue(false)

	assert.Equal(t, 3, queue.Dequeue().(int))
	assert.Equal(t, "strings", queue.Dequeue().(string))
	assert.Equal(t, false, queue.Dequeue().(bool))
	assert.Nil(t, queue.Dequeue())
}
