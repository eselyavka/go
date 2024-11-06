package solutions

import "math"

func canSortArray(nums []int) bool {
	set_bits := countSetBits(nums[0])

	curr_max := nums[0]
	curr_min := nums[0]

	prev_max := math.MinInt

	for i := 1; i < len(nums); i++ {
		curr_set_bits := countSetBits(nums[i])
		if set_bits == curr_set_bits {
			curr_max = MaxInts([]int{curr_max, nums[i]})
			curr_min = MinInts([]int{curr_min, nums[i]})
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
