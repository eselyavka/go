package solutions

import "math"

func minimizedMaximum(n int, quantities []int) int {
	left := 1
	right := MaxInts(quantities)

	for left < right {
		mid := left + (right-left)/2
		target := 0
		for _, num := range quantities {
			target += int(math.Ceil(float64(num) / float64(mid)))
		}
		if target > n {
			left = mid + 1
		} else {
			right = mid
		}
	}

	return left
}
