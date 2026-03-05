package solutions

func coinChange(coins []int, amount int) int {
	dp := make([]int, amount+1, amount+1)
	for i := 1; i < amount+1; i++ {
		dp[i] = 10001
	}

	for _, coin := range coins {
		for j := coin; j < amount+1; j++ {
			dp[j] = MinInts([]int{dp[j], dp[j-coin] + 1})
		}
	}

	if dp[amount] == 10001 {
		return -1
	}

	return dp[amount]
}
