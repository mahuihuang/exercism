package atbash

import (
	"strings"
	"unicode"
)

func Atbash(s string) string {
	output := []rune{}
	s = strings.ToLower(s)
	for i, count := 0, 0; i < len(s); i++ {
		char := rune(s[i])
		if unicode.IsLetter(char) || unicode.IsDigit(char) {
			if unicode.IsLetter(char) {
				output = append(output, 'z'-(char-'a'))
			} else {
				output = append(output, char)
			}
			count++
			if count == 5 {
				output = append(output, ' ')
				count = 0
			}
		}
	}

	return strings.TrimSpace(string(output))
}
