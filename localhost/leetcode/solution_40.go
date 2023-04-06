package solutions

import "sort"

func backtracking_40(local_ans, candidates []int, pos, remain int, ans *[][]int) {
	if remain < 0 {
		return
	}

	if remain == 0 {
		buf := make([]int, len(local_ans))
		copy(buf, local_ans)
		*ans = append(*ans, buf)
		return
	}

	prev := -1
	for i := pos; i < len(candidates); i++ {
		if prev == candidates[i] {
			continue
		}

		local_ans = append(local_ans, candidates[i])
		backtracking_40(local_ans, candidates, i+1, remain-candidates[i], ans)
		local_ans = local_ans[:len(local_ans)-1]

		prev = candidates[i]
	}
}

func combinationSum2(candidates []int, target int) [][]int {
	local_ans := make([]int, 0)
	ans := make([][]int, 0)

	sort.Ints(candidates)

	backtracking_40(local_ans, candidates, 0, target, &ans)

	return ans
}
