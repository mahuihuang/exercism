package wordcount

import (
	"regexp"
	"strings"
)

type Frequency map[string]int

func WordCount(phrase string) Frequency {
	words := make(Frequency)

	for _, v := range regexp.MustCompile(`\w+('\w+)?`).FindAllString(phrase, -1) {
		words[strings.ToLower(v)]++
	}
	return words
}
