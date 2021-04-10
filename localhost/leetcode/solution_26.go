package solutions

func removeDuplicates(nums []int) int {
	if len(nums) == 0 {
		return 0
	}

	if len(nums) == 1 {
		return 1
	}

	n := len(nums)
	j := 0
	i := 1
	for i = 0; i < n-1; i++ {
		if nums[i] != nums[i+1] {
			nums[j] = nums[i]
			j++
		}
	}

	nums[j] = nums[n-1]
	j++

	return j
}
