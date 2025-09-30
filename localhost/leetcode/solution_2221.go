package solutions

func triangularSum(nums []int) int {
	if len(nums) == 1 {
		return nums[0]
	}

	for len(nums) > 1 {
		n := len(nums)
		for i := 1; i < n; i++ {
			nums[i-1] = (nums[i] + nums[i-1]) % 10
		}
		nums = nums[:n-1]
	}

	return nums[0]
}
