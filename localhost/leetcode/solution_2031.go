package solutions

func subarraysWithMoreZerosThanOnes(nums []int) int {
	ans := 0
	dp := 0
	sum := 0

	counter := make(map[int]int)
	counter[0] = 1

	for _, num := range nums {
		if num == 1 {
			sum++
			dp += counter[sum-1]
		} else {
			sum--
			dp -= counter[sum]
		}
		ans += dp
		counter[sum]++
	}

	return ans % 1000000007
}
