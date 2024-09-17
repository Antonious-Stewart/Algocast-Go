package binary_trees

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestPreOrderSearch(t *testing.T) {
	printedTree := PreOrderSearch(&MockRoot)

	assert.Equal(t, []int{7, 23, 5, 4, 3, 18, 21}, printedTree)
}
