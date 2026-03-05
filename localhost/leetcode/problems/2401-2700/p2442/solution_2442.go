package p2442

import "localhost/leetcode/util"

func countDistinctIntegers(nums []int) int {
	ans := make(map[int]struct{})

	for i := 0; i < len(nums); i++ {
		ans[nums[i]] = struct{}{}
		ans[util.ReverseInt(nums[i])] = struct{}{}
	}

	return len(ans)
}
