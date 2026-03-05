package solutions

func maxSubArray(nums []int) int {
	ans := MinInt
	acc := 0

	for _, num := range nums {
		acc += num
		ans = MaxInts([]int{ans, acc})
		acc = MaxInts([]int{acc, 0})
	}
	return ans
}
