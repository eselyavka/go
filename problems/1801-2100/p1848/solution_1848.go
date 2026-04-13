package p1848

func getMinDistance(nums []int, target int, start int) int {
	n := len(nums)

	for d := 0; d < n; d++ {
		if start-d >= 0 && nums[start-d] == target {
			return d
		}
		if d != 0 && start+d < n && nums[start+d] == target {
			return d
		}
	}

	return -1
}
