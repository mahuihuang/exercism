package bottlesong

import (
	"fmt"
	"strings"
)

var numberTranslate = []string{"no", "One", "Two", "Three", "Four", "Five", "Six", "Seven", "Eight", "Nine", "Ten"}

func Recite(startBottles, takeDown int) []string {
	output := []string{}
	if takeDown > startBottles {
		return nil
	}
	for i := 0; i < takeDown; i++ {
		order := startBottles - i
		statement := ""
		//first and second statement
		if order == 1 {
			statement = fmt.Sprintf("%s green bottle hanging on the wall,", numberTranslate[order])
		} else {
			statement = fmt.Sprintf("%s green bottles hanging on the wall,", numberTranslate[order])
		}
		output = append(output, statement, statement, "And if one green bottle should accidentally fall,")

		order--
		// last statement
		if order == 1 {
			statement = fmt.Sprintf("There'll be %s green bottle hanging on the wall.", strings.ToLower(numberTranslate[order]))
		} else {
			statement = fmt.Sprintf("There'll be %s green bottles hanging on the wall.", strings.ToLower(numberTranslate[order]))
		}
		output = append(output, statement, "")
	}
	if output[len(output)-1] == "" {
		output = output[:len(output)-1]
	}
	return output
}
