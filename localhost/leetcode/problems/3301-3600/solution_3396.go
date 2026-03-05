package solutions

func minimumOperations_3396(nums []int) int {
	n := len(nums)
	operations := 0
	i := 0

	for i < n {
		seen := make(map[int]bool)
		dupFound := false

		for j := i; j < n; j++ {
			if seen[nums[j]] {
				dupFound = true
				break
			}
			seen[nums[j]] = true
		}

		if !dupFound {
			return operations
		}
		i += 3
		operations++
	}

	return operations
}
