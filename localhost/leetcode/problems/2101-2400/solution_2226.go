package solutions

func maximumCandies(candies []int, k int64) int {
	k_ := int(k)
	s := sum(candies)

	if s < k_ {
		return 0
	}

	if s == k_ {
		return 1
	}

	if k == 1 {
		return MaxInts(candies)
	}

	l := 0
	r := s / 2

	for l < r {
		mid := (l + r + 1) / 2
		arr := make([]int, len(candies))
		for i := 0; i < len(candies); i++ {
			arr[i] = candies[i] / mid
		}

		if k_ > sum(arr) {
			r = mid - 1
		} else {
			l = mid
		}
	}

	return l
}
