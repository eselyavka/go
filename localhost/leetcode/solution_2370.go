package solutions

func inRange(dp []int, l, r int) int {
	cnt := 0
	for i := l; i <= r; i++ {
		cnt = MaxInts([]int{cnt, dp[i]})
	}
	return cnt
}

func longestIdealString(s string, k int) int {
	n := len(s)
	if n == 1 {
		return 1
	}

	dp := make([]int, 26, 26)
	for _, c := range s {
		idx := c - []rune("a")[0]
		dp[idx] = inRange(dp, MaxInts([]int{0, int(idx) - k}), MinInts([]int{25, int(idx) + k})) + 1
	}

	return MaxInts(dp)
}
