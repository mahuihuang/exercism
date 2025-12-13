package prime

import (
	"errors"
	"math/big"
)

// Nth returns the nth prime number. An error must be returned if the nth prime number can't be calculated ('n' is equal or less than zero)
func Nth(n int) (int, error) {
	if n < 1 {
		return 0, errors.New("the nth prime number can't be calculated")
	}

	for i, count := 2, 0; ; i++ {
		if big.NewInt(int64(i)).ProbablyPrime(1) {
			count++
		}
		if count == n {
			return i, nil
		}
	}
}
