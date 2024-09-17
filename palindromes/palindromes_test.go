package palindromes

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestSolution(t *testing.T) {
	input := "racecar"
	output := Solution(input)

	assert.Equal(t, output, true, "Racecar is a palindrome")
}

func TestEmptySolution(t *testing.T) {
	input := ""
	output := Solution(input)

	assert.Equal(t, output, false, "Empty string returns false")
}

func TestNotAPalindrome(t *testing.T) {
	input := "apple"
	output := Solution(input)

	assert.Equal(t, output, false)
}
