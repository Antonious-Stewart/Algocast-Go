package quick_sort

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

type solution struct {
	arg1       []int
	arg2, arg3 int
}

var inputs = [][]int{
	{12, 7, 5, 23, 15},
	{34, 18, 22, 47, 9},
	{25, 11, 9, 31, 42},
	{8, 14, 6, 19, 3},
}

var expected = [][]int{
	{5, 7, 12, 15, 23},
	{9, 18, 22, 34, 47},
	{9, 11, 25, 31, 42},
	{3, 6, 8, 14, 19},
}

var solutions = []solution{
	{inputs[0], 0, len(inputs[0]) - 1},
	{inputs[1], 0, len(inputs[1]) - 1},
	{inputs[2], 0, len(inputs[2]) - 1},
	{inputs[3], 0, len(inputs[3]) - 1},
}

func TestSolution(t *testing.T) {
	for i, test := range solutions {
		Solution(test.arg1, test.arg2, test.arg3)
		assert.Exactly(t, test.arg1, expected[i], "The input of %v matches %v", test.arg1, expected[i])
	}
}
