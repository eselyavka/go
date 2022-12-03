package solutions

func countDistinctIntegers(nums []int) int {
	ans := make(map[int]struct{})

	for i := 0; i < len(nums); i++ {
		ans[nums[i]] = struct{}{}
		ans[reverseInt(nums[i])] = struct{}{}
	}

	return len(ans)
}
