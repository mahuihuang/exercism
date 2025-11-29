package strain

// Implement the "Keep" and "Discard" function in this file.

// You will need typed parameters (aka "Generics") to solve this exercise.
// They are not part of the Exercism syllabus yet but you can learn about
// them here: https://go.dev/tour/generics/1

func Keep[T string | int | [][]int | []int](collection []T, filter func(T) bool) []T {
	c := make([]T, 0)
	for _, v := range collection {
		if filter(v) {
			c = append(c, v)
		}
	}
	return c
}

func Discard[T string | int | [][]int | []int](collection []T, filter func(T) bool) []T {
	c := make([]T, 0)
	for _, v := range collection {
		if !filter(v) {
			c = append(c, v)
		}
	}
	return c
}
