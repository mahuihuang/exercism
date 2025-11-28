package isogram

import (
	"strings"
	"unicode"
)

func IsIsogram(word string) bool {
	letters := make(map[rune]bool)
	for _, w := range strings.ToLower(word) {
		if !unicode.IsLetter(w) {
			continue
		}
		if letters[w] {
			return false
		}
		letters[w] = true
	}
	return true
}
