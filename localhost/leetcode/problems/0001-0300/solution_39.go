package solutions

func rec_39(target, idx, sum int, local_ans, candidates []int, ans *[][]int) {
	if idx >= len(candidates) || sum > target {
		return
	}

	if sum == target {
		copy_local_ans := make([]int, len(local_ans))
		copy(copy_local_ans, local_ans)
		*ans = append(*ans, copy_local_ans)

		return
	}

	local_ans = append(local_ans, candidates[idx])

	rec_39(target, idx, sum+candidates[idx], local_ans, candidates, ans)

	local_ans = local_ans[:len(local_ans)-1]

	rec_39(target, idx+1, sum, local_ans, candidates, ans)
}

func combinationSum(candidates []int, target int) [][]int {
	var ans [][]int
	local_ans := make([]int, 0)

	rec_39(target, 0, 0, local_ans, candidates, &ans)

	return ans
}
