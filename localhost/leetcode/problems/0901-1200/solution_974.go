package solutions

func subarraysDivByK(nums []int, k int) int {
	if len(nums) == 1 {
		if nums[0]%k == 0 {
			return 1
		}
		return 0
	}

	ans := 0
	running_sum := 0

	c := make(map[int]int)
	c[0] = 1

	for _, num := range nums {
		running_sum += num
		mod_k := running_sum % k
		if mod_k < 0 {
			mod_k += k
		}
		ans += c[mod_k]
		c[mod_k]++

	}
	return ans
}
