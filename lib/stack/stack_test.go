package stack_test

import (
	"github.com/qulia/go-qulia/lib/stack"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestStackBasic(t *testing.T) {
	testStack := stack.Stack{}
	verify(t, &testStack)
}

func verify(t *testing.T, stack stack.Interface) {
	stack.Push(3)
	stack.Push("strings")
	stack.Push(false)

	assert.Equal(t, false, stack.Pop().(bool))
	assert.Equal(t, "strings", stack.Pop().(string))
	assert.Equal(t, 3, stack.Pop().(int))
	assert.Nil(t, stack.Pop())
}