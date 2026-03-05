package solutions

func minimumSplits(nums []int) int {
	n := len(nums)
	ans := 1
	curr := nums[0]

	for i := 1; i < n; i++ {
		curr = GCD(curr, nums[i])
		if curr == 1 {
			curr = nums[i]
			ans++
		}
	}
	return ans
}
