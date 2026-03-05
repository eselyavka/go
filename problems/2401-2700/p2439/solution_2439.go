package p2439

import "github.com/eseliavka/go/util"

func minimizeArrayValue(nums []int) int {
	ans := 0
	prefixSum := 0

	for idx, num := range nums {
		prefixSum += num
		ans = util.MaxInts([]int{ans, (prefixSum + idx) / (idx + 1)})
	}

	return ans
}
