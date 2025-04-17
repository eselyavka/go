package solutions

import "math"

func binpow(x, y, mod int64) int64 {
	if y == 0 {
		return 1
	}

	res := binpow(x, y/2, mod)

	if y%2 == 1 {
		return (res * res * x) % mod
	}

	return (res * res) % mod
}

func countGoodNumbers(n int64) int {
	mod := int64(math.Pow(10, 9) + 7)

	return int((binpow(5, (n+1)/2, mod) * binpow(4, n/2, mod)) % mod)
}
