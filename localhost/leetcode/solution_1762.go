package solutions

import "sort"

func findBuildings(heights []int) []int {
	if len(heights) == 1 {
		return []int{0}
	}

	r := len(heights) - 2
	res := []int{r + 1}
	max_so_far := heights[r+1]

	for r >= 0 {
		if heights[r] > heights[r+1] && heights[r] > max_so_far {
			max_so_far = heights[r]
			res = append(res, r)
		}
		r -= 1
	}

	sort.Slice(res, func(i, j int) bool { return res[i] < res[j] })

	return res
}
