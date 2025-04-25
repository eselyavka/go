package solutions

func countInterestingSubarrays(nums []int, modulo int, k int) int64 {
	n := len(nums)
	m := make(map[int]int, 0)
	m[0] = 1
	ans := int64(0)
	prefix := 0

	for i := 0; i < n; i++ {
		if nums[i]%modulo == k {
			prefix++
		}
		ans += int64(m[(prefix+modulo-k)%modulo])
		m[prefix%modulo]++
	}

	return ans
}
