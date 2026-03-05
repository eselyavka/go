package solutions

func maxScore(s string) int {
	n := len(s)
	ans := -1
	oneCnt := 0
	zeroCnt := 0
	for i := 0; i < n; i++ {
		if s[i] == '1' {
			oneCnt++
		}
	}
	for i := 0; i < n-1; i++ {
		if s[i] == '0' {
			zeroCnt++
		} else {
			oneCnt--
		}

		ans = MaxInts([]int{ans, zeroCnt + oneCnt})
	}

	return ans
}
