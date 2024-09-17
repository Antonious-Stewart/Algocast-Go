package palindromes

import "strings"

func Solution(str string) bool {
	if str == "" {
		return false
	}

	normalize := strings.ToLower(str)
	n := len(normalize)

	for i := 0; i < n/2; i++ {
		if normalize[i] != normalize[n-i-1] {
			return false
		}
	}

	return true
}
