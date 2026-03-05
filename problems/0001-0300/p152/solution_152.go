package p152

import "github.com/eseliavka/go/util"

func maxProduct(nums []int) int {
	ans := nums[0]
	currMin := 1
	currMax := 1

	for _, num := range nums {
		productMax := num * currMax
		productMin := num * currMin
		currMax = util.MaxInts([]int{productMax, productMin, num})
		currMin = util.MinInts([]int{productMax, productMin, num})
		ans = util.MaxInts([]int{ans, currMax})
	}

	return ans
}
