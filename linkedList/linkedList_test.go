package linked_list

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNodeCreationError(t *testing.T) {
	_, err := NewNode(nil)
	if assert.Error(t, err) {
		assert.Equal(t, "invalid value", err.Error())
	}
}

func TestNodeCreationSuccess(t *testing.T) {
	node, _ := NewNode(1)

	assert.NotNil(t, node, "Expect the node to be defined")
	assert.Nil(t, node.next, "Expect the next value to be nil")
	assert.Equal(t, 1, node.data, "Expect node data to be equal to 1")
}

func TestListCreationFail(t *testing.T) {
	list := List{}
	err := list.NewList(nil)

	if assert.Error(t, err) {
		assert.Equal(t, "invalid value", err.Error())
	}
}

func TestListCreationSuccess(t *testing.T) {
	list := List{}
	list.NewList(1)

	assert.NotNil(t, list.Head, "Should have a valid head node")
	assert.Equal(t, 1, list.Head.data, "Should have a valid head node value")
	assert.Equal(t, list.Head.data, list.Tail.data, "Should have a valid head node value")
	assert.Equal(t, list.Length, 1, "Should have a length of 1")
}

func TestGetLength(t *testing.T) {
	list := List{}

	list.NewList(1)

	assert.Equal(t, 1, list.GetLength(), "Should get the length of the list")
}

func TestGetHead(t *testing.T) {
	list := List{}
	head := list.GetHead()

	assert.Nil(t, head)
}

func TestGetTail(t *testing.T) {
	list := List{}

	list.NewList(1)
	tail := list.GetTail()

	assert.NotNil(t, tail)
}
