package solutions

func maxProduct(nums []int) int {
	ans := nums[0]
	currMin := 1
	currMax := 1

	for _, num := range nums {
		productMax := num * currMax
		productMin := num * currMin
		currMax = MaxInts([]int{productMax, productMin, num})
		currMin = MinInts([]int{productMax, productMin, num})
		ans = MaxInts([]int{ans, currMax})
	}

	return ans
}
