package solutions

func rob(nums []int) int {
	if len(nums) == 1 {
		return nums[0]
	}

	if len(nums) == 2 {
		return MaxInts(nums)
	}

	dp := make([]int, len(nums))
	dp[0] = nums[0]

	for i := 1; i < len(nums); i++ {
		num := nums[i]
		if i >= 2 {
			num += dp[i-2]
		}
		if num > dp[i-1] {
			dp[i] = num
		} else {
			dp[i] = dp[i-1]
		}
	}

	return dp[len(dp)-1]
}
