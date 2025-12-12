package anagram

import (
	"sort"
	"strings"
)

func Detect(subject string, candidates []string) []string {
	output := []string{}
	sortSubject := sortString(strings.ToLower(subject))

	for _, v := range candidates {
		if strings.EqualFold(v, subject) {
			continue
		}
		if sortSubject == sortString(strings.ToLower(v)) {
			output = append(output, v)
		}
	}
	return output
}

func sortString(s string) string {
	chars := strings.Split(s, "")
	sort.Strings(chars)
	return strings.Join(chars, "")
}
