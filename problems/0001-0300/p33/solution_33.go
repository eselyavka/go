package p33

import "github.com/eseliavka/go/util"

func search(nums []int, target int) int {
	n := len(nums)
	if n == 1 {
		if nums[0] == target {
			return 0
		} else {
			return -1
		}
	}

	ans := util.BinarySearch(0, n-1, nums, target)

	return ans
}
