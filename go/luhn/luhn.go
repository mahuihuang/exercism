package luhn

import (
	"strconv"
	"strings"
)

// 0+1+5+8+4+8+2+7+5
func Valid(id string) bool {
	trimId := strings.ReplaceAll(id, " ", "")
	if len(trimId) <= 1 {
		return false
	}
	sum := 0
	for i, isEven := len(trimId)-1, false; i >= 0; i-- {
		n, err := strconv.Atoi(string(trimId[i]))
		if err != nil {
			return false
		}
		if m := n * 2; isEven {
			sum += m
			if m > 9 {
				sum -= 9
			}
		} else {
			sum += n
		}
		isEven = !isEven
	}
	return sum%10 == 0
}
