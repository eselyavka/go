package solutions

func minimizeArrayValue(nums []int) int {
	ans := 0
	prefix_sum := 0

	for idx, num := range nums {
		prefix_sum += num
		ans = MaxInts([]int{ans, (prefix_sum + idx) / (idx + 1)})
	}

	return ans
}
