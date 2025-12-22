package armstrong

import "math"

func IsNumber(n int) bool {
	len := 0
	for number := n; number > 0; number /= 10 {
		len++
	}

	var sum float64 = 0
	for number := n; number > 0; number /= 10 {
		sum += math.Pow(float64(number%10), float64(len))
	}

	return int(sum) == n
}
