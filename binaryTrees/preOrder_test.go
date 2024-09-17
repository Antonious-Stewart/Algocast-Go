package binary_trees

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

var MockRoot = Node{
	Data: 7,
	Left: &Node{
		Data: 23,
		Left: &Node{
			Data: 5,
		},
		Right: &Node{
			Data: 4,
		},
	},
	Right: &Node{
		Data: 3,
		Left: &Node{
			Data: 18,
		},
		Right: &Node{
			Data: 21,
		},
	},
}

func TestPreOrderSearch(t *testing.T) {
	printedTree := PreOrderSearch(&root)

	assert.Equal(t, []int{7, 23, 5, 4, 3, 18, 21}, printedTree)
}
