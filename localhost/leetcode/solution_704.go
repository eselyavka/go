package solutions

func search_704(nums []int, target int) int {
	n := len(nums)

	if n == 1 {
		if nums[0] == target {
			return 0
		} else {
			return -1
		}
	}

	ans := binarySearch(0, n-1, nums, target)

	return ans
}
