package sieve

import "math/big"

func Sieve(limit int) []int {
	primes := []int{}
	for number := range limit + 1 {
		if big.NewInt(int64(number)).ProbablyPrime(1) {
			primes = append(primes, number)
		}
	}
	return primes
}
