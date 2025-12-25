package series

func All(n int, s string) []string {
	if n > len(s) {
		return nil
	}
	series := []string{}
	for i := 0; i+n <= len(s); i++ {
		series = append(series, s[i:i+n])
	}
	return series
}

func UnsafeFirst(n int, s string) string {
	return All(n, s)[0]
}

func First(n int, s string) (first string, ok bool) {
	if n > len(s) {
		return "", false
	} else {
		return UnsafeFirst(n, s), true
	}
}
