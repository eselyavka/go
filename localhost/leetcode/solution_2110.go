package solutions

func getDescentPeriods(prices []int) int64 {
	n := len(prices)
	ans := 0
	run := 0

	for i := 0; i < n; i++ {
		if i > 0 && prices[i-1]-prices[i] == 1 {
			run++
		} else {
			run = 1
		}
		ans += run
	}

	return int64(ans)
}
