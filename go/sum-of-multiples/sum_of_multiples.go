package summultiples

import (
	"slices"
)

func SumMultiples(limit int, divisors ...int) int {
	sum := 0
	items := []int{}
	for _, divisor := range divisors {
		if divisor >= limit || divisor == 0 {
			continue
		}
		for i := 1; ; i++ {
			if v := divisor * i; v < limit {
				items = append(items, v)
			} else {
				break
			}
		}
	}
	slices.Sort(items)
	// slices.Compact 移除相邻重复的元素
	for _, v := range slices.Compact(items) {
		sum += v
	}
	return sum
}
