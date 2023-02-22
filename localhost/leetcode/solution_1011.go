package solutions

func binary_search_1011(target int, weights []int, days int) bool {
	cnt := 1
	sub_sub := 0

	for _, weight := range weights {
		if weight > target {
			return false
		}
		sub_sub += weight

		if sub_sub > target {
			cnt++
			sub_sub = weight
		}
	}

	if cnt <= days {
		return true
	}

	return false
}

func shipWithinDays(weights []int, days int) int {
	ans := 0
	minw := MaxInts(weights)
	maxw := sum(weights)

	for minw <= maxw {
		mid := (minw + maxw) / 2
		if binary_search_1011(mid, weights, days) {
			ans = mid
			maxw = mid - 1
		} else {
			minw = mid + 1
		}
	}

	return ans
}
