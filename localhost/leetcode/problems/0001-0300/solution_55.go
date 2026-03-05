package solutions

func canJump(nums []int) bool {
	n := len(nums)

	if n == 1 {
		return true
	}

	dp := make([]bool, n, n)

	goal := n - 1

	for i := n - 1; i > -1; i-- {
		if i+nums[i] >= goal {
			goal = i
			dp[i] = true
		}
	}

	return dp[0]
}
