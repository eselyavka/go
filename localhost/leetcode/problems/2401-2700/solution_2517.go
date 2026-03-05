package solutions

import "sort"

func check(array []int, target, k int) bool {
	prev := array[0]
	cnt := 1
	for i := 1; i < len(array); i++ {
		diff := array[i] - prev
		if diff >= target {
			prev = array[i]
			cnt++
		}
		if cnt == k {
			return true
		}
	}

	return cnt == k
}
func maximumTastiness(price []int, k int) int {
	sort.Ints(price)
	if k == 2 {
		return price[len(price)-1] - price[0]
	}

	low := 0
	high := price[len(price)-1] - price[0]

	ans := -1
	for low <= high {
		mid := (low + high) / 2
		if check(price, mid, k) {
			ans = mid
			low = mid + 1
		} else {
			high = mid - 1
		}
	}

	return ans
}
