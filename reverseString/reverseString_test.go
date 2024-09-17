package reverse_string

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestReverseString(t *testing.T) {
	input := "Hello World"
	output := Solution(input)

	assert.Equal(t, output, "dlroW olleH")
}
