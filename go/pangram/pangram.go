package pangram

import (
	"strings"
)

func IsPangram(input string) bool {
	lower := strings.ToLower(input)
	for i := 'a'; i <= 'z'; i++ {
		if !strings.ContainsRune(lower, i) {
			return false
		}
	}
	return true
}
