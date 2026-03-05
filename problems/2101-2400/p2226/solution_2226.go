package p2226

import "github.com/eseliavka/go/util"

func maximumCandies(candies []int, k int64) int {
	k_ := int(k)
	s := util.Sum(candies)

	if s < k_ {
		return 0
	}

	if s == k_ {
		return 1
	}

	if k == 1 {
		return util.MaxInts(candies)
	}

	l := 0
	r := s / 2

	for l < r {
		mid := (l + r + 1) / 2
		arr := make([]int, len(candies))
		for i := 0; i < len(candies); i++ {
			arr[i] = candies[i] / mid
		}

		if k_ > util.Sum(arr) {
			r = mid - 1
		} else {
			l = mid
		}
	}

	return l
}
