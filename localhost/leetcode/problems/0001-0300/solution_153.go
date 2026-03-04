package solutions

func findMin(nums []int) int {
	n := len(nums)

	if n == 1 {
		return nums[0]
	}

	return binarySearchMin(0, n-1, nums)
}
