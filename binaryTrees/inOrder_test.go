package binary_trees

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestIOnOrderSearch(t *testing.T) {
	printedTree := InOrderSearch(&MockRoot)

	assert.Equal(t, []int{5, 23, 4, 7, 18, 3, 21}, printedTree)
}
