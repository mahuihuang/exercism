package rotationalcipher

import (
	"unicode"
)

func RotationalCipher(plain string, shiftKey int) string {
	var cipher []rune
	for _, char := range plain {
		if unicode.IsLetter(char) {
			baseChar := 0
			if unicode.IsLower(char) {
				baseChar = 'a'
			} else {
				baseChar = 'A'
			}
			shiftChar := (shiftKey+int(char)-baseChar)%26 + baseChar
			cipher = append(cipher, rune(shiftChar))
		} else {
			cipher = append(cipher, char)
		}
	}

	return string(cipher)
}
