package solutions

func countPairs(nums []int, k int) int {
	n := len(nums)
	ans := 0
	for i := 0; i < n; i++ {
		for j := i + 1; j < n; j++ {
			if nums[i] == nums[j] && (i*j)%k == 0 {
				ans++
			}
		}
	}

	return ans
}
