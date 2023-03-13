package solutions

func minStartValue(nums []int) int {
	ans := MinInt

	sum_so_far := 0

	for _, num := range nums {
		sum_so_far += num
		if sum_so_far <= 0 {
			x := -1*sum_so_far + 1
			ans = MaxInts([]int{ans, x})
		}
	}

	if ans == MinInt {
		return 1
	}

	return ans
}
