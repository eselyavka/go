package p3507

import "math"

func minimumPairRemoval(nums []int) int {
	ans := 0

	checkIsNoneDecreasing := func(nums []int) bool {
		n := len(nums)
		for i := 1; i < n; i++ {
			if nums[i] < nums[i-1] {
				return false
			}
		}

		return true
	}

	for !checkIsNoneDecreasing(nums) && len(nums) > 1 {
		min := math.MaxInt
		idxs := make([]int, 2)
		for i := 1; i < len(nums); i++ {
			s := nums[i] + nums[i-1]
			if s < min {
				idxs[0] = i - 1
				idxs[1] = i
				min = s
			}
		}
		x := idxs[0]
		y := idxs[1]

		buf := make([]int, 0, len(nums)-1)
		buf = append(buf, nums[:x]...)
		buf = append(buf, nums[x]+nums[y])
		buf = append(buf, nums[y+1:]...)
		nums = buf
		ans++
	}

	return ans
}
