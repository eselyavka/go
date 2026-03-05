package p3011

import (
	"localhost/leetcode/util"
	"math"
)

func canSortArray(nums []int) bool {
	set_bits := util.CountSetBits(nums[0])

	curr_max := nums[0]
	curr_min := nums[0]

	prev_max := math.MinInt

	for i := 1; i < len(nums); i++ {
		curr_set_bits := util.CountSetBits(nums[i])
		if set_bits == curr_set_bits {
			curr_max = util.MaxInts([]int{curr_max, nums[i]})
			curr_min = util.MinInts([]int{curr_min, nums[i]})
		} else {
			if curr_min < prev_max {
				return false
			}

			prev_max = curr_max

			curr_max = nums[i]
			curr_min = nums[i]

			set_bits = curr_set_bits
		}
	}
	if curr_min < prev_max {
		return false
	}

	return true
}
