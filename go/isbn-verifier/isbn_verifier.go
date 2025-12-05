package isbn

import (
	"strconv"
	"unicode"
)

func IsValidISBN(isbn string) bool {
	d := 10
	sum := 0
	for i := 0; i < len(isbn); i++ {
		v := isbn[i]
		switch {
		case unicode.IsDigit(rune(v)):
			number, _ := strconv.Atoi(string(v))
			sum += number * d
			d--
		case v == 'X':
			if i != len(isbn)-1 {
				return false
			}
			sum += 10 * d
			d--
		case v == '-':
			continue
		default:
			return false
		}
	}
	if d != 0 {
		return false
	}
	return sum%11 == 0
}
