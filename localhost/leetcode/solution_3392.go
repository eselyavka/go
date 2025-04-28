package solutions

func countSubarrays(nums []int) int {
	ans := 0
	n := len(nums)

	for i := 0; i < n; i++ {
		if i+2 >= n {
			break
		}
		left := i
		mid := i + 1
		right := i + 2
		if nums[mid]%2 == 0 && (nums[left]+nums[right]) == nums[mid]/2 {
			ans++
		}
	}

	return ans
}
