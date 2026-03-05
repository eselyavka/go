package solutions

func countPrimes(n int) int {
	if n < 2 {
		return 0
	}

	primes := make([]bool, n, n)
	for i := 2; i < n; i++ {
		primes[i] = true
	}

	p := 2

	for p*p < n {
		if primes[p] {
			for i := p * p; i < n; i += p {
				primes[i] = false
			}
		}
		p++
	}

	return nTrue(primes)
}
