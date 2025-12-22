package cipher

import (
	"strings"
	"unicode"
)

type shift int

type vigenere struct {
	key string
}

// Define the shift and vigenere types here.
// Both types should satisfy the Cipher interface.
func NewCaesar() Cipher {
	return NewShift(3)
}

func NewShift(distance int) Cipher {
	if (distance < 26 && distance > 0) || (distance < 0 && distance > -26) {
		return shift(distance)
	}
	return nil
}

func (c shift) Encode(input string) string {
	output := []rune{}
	for _, r := range strings.ToLower(input) {
		if unicode.IsLetter(r) {
			output = append(output, shiftRune(r, c))
		}
	}
	return string(output)
}

func (c shift) Decode(input string) string {
	output := []rune{}
	for _, r := range input {
		output = append(output, shiftRune(r, -c))
	}
	return string(output)
}

func shiftRune(r rune, s shift) rune {
	char := (rune(s)+r-'a')%26 + 'a'
	if char < 'a' {
		char = 'z' - ('a' - char) + 1
	}
	return char
}
func NewVigenere(key string) Cipher {
	isValidKey := false
	for _, char := range key {
		if unicode.IsLetter(char) && unicode.IsLower(char) {
			if char != 'a' {
				isValidKey = true
			}
		} else {
			return nil
		}
	}

	if isValidKey {
		return vigenere{key: key}
	} else {
		return nil
	}
}

func (v vigenere) Encode(input string) string {
	output := ""
	input = strings.ToLower(input)
	position := 0
	for _, char := range input {
		if !unicode.IsLetter(char) {
			continue
		}
		// 取余，保证一直可以循环
		position %= len(v.key)
		cipher := NewShift(int(v.key[position] - 'a'))
		if cipher == nil {
			output += string(char)
		} else {
			output += cipher.Encode(string(char))
		}
		position++
	}
	return output
}
func (v vigenere) Decode(input string) string {
	output := ""
	position := 0
	for _, char := range input {
		position %= len(v.key)
		cipher := NewShift(int(v.key[position] - 'a'))
		if cipher == nil {
			output += string(char)
		} else {
			output += cipher.Decode(string(char))
		}
		position++
	}
	return output
}
