package lsproduct

import (
	"errors"
	"strconv"
)

func LargestSeriesProduct(digits string, span int) (int64, error) {
	var maxProduct int64

	if digits == "" || span > len(digits) || span <= 0 {
		return maxProduct, errors.New("illage digits")
	}

	for start, end := 0, span; end < len(digits)+1; start, end = start+1, end+1 {
		series := digits[start:end]
		if series == "" {
			continue
		}
		number, err := strconv.Atoi(series)
		if err != nil {
			return 0, err
		}
		product := number % 10
		number /= 10
		for range span - 1 {
			product *= number % 10
			number /= 10
		}
		if v := int64(product); v > maxProduct {
			maxProduct = v
		}
	}
	return maxProduct, nil
}
