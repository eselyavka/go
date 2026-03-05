package solutions

import "sort"

func maxSatisfaction(satisfaction []int) int {
	n := len(satisfaction)

	sort.Ints(satisfaction)

	if satisfaction[n-1] <= 0 {
		return 0
	}

	ans := 0

	i := 0
	for i < n {
		local_ans := satisfaction[i]
		j := i + 1
		k := 2
		for j < n {
			local_ans += satisfaction[j] * k
			j++
			k++
		}

		ans = MaxInts([]int{ans, local_ans})
		i++
	}

	return ans
}
