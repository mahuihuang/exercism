package listops

// IntList is an abstraction of a list of integers which we can define methods on
type IntList []int

// Foldl write by AI
func (s IntList) Foldl(fn func(int, int) int, initial int) int {
	acc := initial
	for _, v := range s {
		acc = fn(acc, v)
	}
	return acc
}

// Foldr write by AI
func (s IntList) Foldr(fn func(int, int) int, initial int) int {
	acc := initial
	for i := len(s) - 1; i >= 0; i-- {
		acc = fn(s[i], acc)
	}
	return acc
}

func (s IntList) Filter(fn func(int) bool) IntList {
	output := IntList{}
	for _, v := range s {
		if fn(v) {
			output = append(output, v)
		}
	}
	return output
}

func (s IntList) Length() int {
	return len(s)
}

func (s IntList) Map(fn func(int) int) IntList {
	m := IntList{}
	for _, v := range s {
		m = append(m, fn(v))
	}
	return m
}

func (s IntList) Reverse() IntList {
	for i, j := 0, len(s)-1; i < j; i, j = i+1, j-1 {
		s[i], s[j] = s[j], s[i]
	}
	return s
}

func (s IntList) Append(lst IntList) IntList {
	return append(s, lst...)
}

func (s IntList) Concat(lists []IntList) IntList {
	for _, l := range lists {
		s = append(s, l...)
	}
	return s
}
