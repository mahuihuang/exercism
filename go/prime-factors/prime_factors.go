package prime

// import "math/big"

func Factors(n int64) []int64 {
	primeList := []int64{}
	var prime int64 = 2
	for i := prime; n != 1; i++ {
		for prime = i; n%prime == 0; n /= prime {
			primeList = append(primeList, prime)
		}
	}
	return primeList
}
