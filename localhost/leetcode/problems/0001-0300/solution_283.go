package solutions

func moveZeroes(nums []int) {
	i := 0
	j := 0
	n := len(nums)

	for i < n {
		if nums[i] != 0 {
			tmp := nums[j]
			nums[j] = nums[i]
			nums[i] = tmp
			j++
		}
		i++
	}
}
