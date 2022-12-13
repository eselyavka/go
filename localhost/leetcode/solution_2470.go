package solutions

func subarrayLCM(nums []int, k int) int {
	n := len(nums)
	ans := 0

	var curr int
	for i := 0; i < n; i++ {
		curr = nums[i]
		if curr == k {
			ans++
		}
		for j := i + 1; j < n; j++ {
			curr = LCM(curr, nums[j])
			if curr == k {
				ans++
			}

			if curr > k {
				break
			}
		}
	}
	return ans
}
