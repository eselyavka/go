package solutions

func subarrayGCD(nums []int, k int) int {
	n := len(nums)
	ans := 0
	curr := 0

	for i := 0; i < n; i++ {
		if nums[i] == k {
			ans++
		}
		curr = nums[i]
		for j := i + 1; j < n; j++ {
			curr = GCD(curr, nums[j])
			if curr == k {
				ans++
			}
		}
	}
	return ans
}
