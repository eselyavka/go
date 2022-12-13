package solutions

func lengthOfLIS(nums []int) int {
	n := len(nums)
	if n == 1 {
		return 1
	}

	dp := make([]int, n, n)
	for i := 0; i < n; i++ {
		dp[i] = 1
	}

	for i := 1; i < n; i++ {
		for j := 0; j < i; j++ {
			if nums[i] > nums[j] {
				dp[i] = MaxInts([]int{dp[i], 1 + dp[j]})
			}
		}
	}

	return MaxInts(dp)
}
