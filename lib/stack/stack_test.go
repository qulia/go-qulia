package stack_test

import (
	"testing"

	"github.com/qulia/go-qulia/lib/stack"
	"github.com/stretchr/testify/assert"
)

func TestStackBasic(t *testing.T) {
	testStack := stack.Stack{}
	verify(t, &testStack)
}

func verify(t *testing.T, stack stack.Interface) {
	stack.Push(3)
	stack.Push("strings")
	stack.Push(false)

	assert.Equal(t, false, stack.Peek().(bool))
	assert.Equal(t, false, stack.Pop().(bool))
	assert.Equal(t, "strings", stack.Peek().(string))
	assert.Equal(t, "strings", stack.Pop().(string))
	assert.Equal(t, 3, stack.Peek().(int))
	assert.Equal(t, 3, stack.Pop().(int))
	assert.Nil(t, stack.Peek())
	assert.Nil(t, stack.Pop())
}
