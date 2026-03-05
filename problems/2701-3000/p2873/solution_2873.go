package p2873

import "github.com/eseliavka/go/util"

func maximumTripletValueOne(nums []int) int64 {
	n := len(nums)
	leftMax := make([]int, n, n)
	rightMax := make([]int, n, n)

	for i := 1; i < n; i++ {
		leftMax[i] = util.MaxInts([]int{nums[i-1], leftMax[i-1]})
		rightMax[n-i-1] = util.MaxInts([]int{nums[n-i], rightMax[n-i]})
	}

	ans := 0
	for i := 1; i < n-1; i++ {
		ans = util.MaxInts([]int{ans, (leftMax[i] - nums[i]) * rightMax[i]})
	}

	return int64(ans)
}
