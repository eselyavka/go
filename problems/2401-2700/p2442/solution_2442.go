package p2442

import "github.com/eseliavka/go/util"

func countDistinctIntegers(nums []int) int {
	ans := make(map[int]struct{})

	for i := 0; i < len(nums); i++ {
		ans[nums[i]] = struct{}{}
		ans[util.ReverseInt(nums[i])] = struct{}{}
	}

	return len(ans)
}
