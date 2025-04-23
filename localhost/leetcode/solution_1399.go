package solutions

func countLargestGroup(n int) int {
	digitSums := make(map[int]int)
	currMax := 0

	for i := 1; i <= n; i++ {
		s := sumDigits(i)
		digitSums[s]++
		currMax = MaxInts([]int{currMax, digitSums[s]})
	}

	ans := 0
	for _, count := range digitSums {
		if count == currMax {
			ans++
		}
	}
	return ans
}
