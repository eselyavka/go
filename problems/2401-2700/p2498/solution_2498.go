package p2498

import "github.com/eseliavka/go/util"

func maxJump(stones []int) int {
	n := len(stones)

	if n == 2 {
		return stones[1]
	}

	ans := util.MinInt
	var i int
	i = 2
	for i < n {
		ans = util.MaxInts([]int{ans, util.Abs(stones[i-2] - stones[i])})
		i += 2
	}

	i = 3
	for i < n {
		ans = util.MaxInts([]int{ans, util.Abs(stones[i-2] - stones[i])})
		i += 2
	}

	return ans
}
