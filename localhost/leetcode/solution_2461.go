package solutions

func maximumSubarraySum(nums []int, k int) int64 {
	n := len(nums)
	ans := 0
	curr_sum := 0
	left := 0
	right := 0
	idx_map := make(map[int]int)

	for idx := 0; idx < n; idx++ {
		curr_num := nums[idx]
		seen := -1
		if v, ok := idx_map[curr_num]; ok {
			seen = v
		}
		for left <= seen || right-left+1 > k {
			curr_sum -= nums[left]
			left++
		}
		idx_map[curr_num] = idx
		curr_sum += nums[right]
		if right-left+1 == k {
			ans = MaxInts([]int{ans, curr_sum})
		}
		right++
	}

	return int64(ans)
}
