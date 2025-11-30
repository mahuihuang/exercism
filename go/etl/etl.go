package etl

import "strings"

func Transform(in map[int][]string) map[string]int {
	output := map[string]int{}
	for point, letters := range in {
		for _, char := range letters {
			output[strings.ToLower(char)] = point
		}
	}
	return output
}
