package solutions

import "sort"

func hIndex(citations []int) int {

	sort.Sort(sort.Reverse(sort.IntSlice(citations)))

	ans := 0

	for idx, num := range citations {
		if num > idx {
			ans++
		}
	}

	return ans
}
