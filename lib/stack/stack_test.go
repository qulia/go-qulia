package stack_test

import (
	"testing"

	"github.com/qulia/go-qulia/lib/stack"
	"github.com/stretchr/testify/assert"
)

func TestStackBasic(t *testing.T) {
	testStack := stack.NewStack[string]()
	verify(t, testStack)
}

func verify(t *testing.T, stack stack.Stack[string]) {
	stack.Push("one")
	stack.Push("two")
	stack.Push("three")

	assert.Equal(t, "three", stack.Peek())
	assert.Equal(t, "three", stack.Pop())
	assert.Equal(t, "two", stack.Peek())
	assert.Equal(t, "two", stack.Pop())
	assert.Equal(t, "one", stack.Peek())
	assert.Equal(t, "one", stack.Pop())
	assert.True(t, stack.IsEmpty())
	assert.Panics(t, func() { stack.Peek() })
	assert.Panics(t, func() { stack.Peek() })
}
