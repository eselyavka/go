package solutions

import "math"

func sumFourDivisors(nums []int) int {
	ans := 0

	for _, n := range nums {
		cnt := 2
		s := 1 + n

		limit := int(math.Sqrt(float64(n)))

		for d := 2; d <= limit; d++ {
			if n%d == 0 {
				other := n / d
				if other == d {
					cnt++
					s += d
				} else {
					cnt += 2
					s += d + other
				}
				if cnt > 4 {
					break
				}
			}
		}

		if cnt == 4 {
			ans += s
		}
	}

	return ans
}
