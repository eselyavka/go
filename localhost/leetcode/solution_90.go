package solutions

import "sort"

func backtrack_90(i int, nums, local_ans []int, ans *[][]int) {
	if i >= len(nums) {
		tmp := make([]int, len(local_ans))
		copy(tmp, local_ans)
		*ans = append(*ans, tmp)
		return
	}

	local_ans = append(local_ans, nums[i])

	backtrack_90(i+1, nums, local_ans, ans)

	local_ans = local_ans[:len(local_ans)-1]

	for i+1 < len(nums) && nums[i+1] == nums[i] {
		i++
	}

	backtrack_90(i+1, nums, local_ans, ans)
}

func subsetsWithDup(nums []int) [][]int {
	ans := make([][]int, 0)

	sort.Ints(nums)

	backtrack_90(0, nums, []int{}, &ans)

	return ans
}
