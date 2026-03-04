package solutions

func countCompleteSubarrays(nums []int) int {
	set := map[int]struct{}{}
	for _, n := range nums {
		set[n] = struct{}{}
	}
	k := len(set)
	n := len(nums)
	ans := 0

	for i := 0; i < n-k+1; i++ {
		set = map[int]struct{}{}
		set[nums[i]] = struct{}{}
		for j := i + 1; j < n; j++ {
			if len(set) == k {
				ans++
			}
			set[nums[j]] = struct{}{}
		}
		if len(set) == k {
			ans++
		}
	}

	return ans
}
