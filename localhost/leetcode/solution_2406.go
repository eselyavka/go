package solutions

import "sort"

func minGroups(intervals [][]int) int {
	n := len(intervals)

	start_points := make([]int, 0)
	end_points := make([]int, 0)

	for _, item := range intervals {
		start := item[0]
		end := item[1]
		start_points = append(start_points, start)
		end_points = append(end_points, end)
	}

	sort.Ints(start_points)
	sort.Ints(end_points)

	l := 0
	r := 0

	ans := 0
	count := 0

	for l < n && r < n {
		s := start_points[l]
		e := end_points[r]
		if s <= e {
			l++
			count++
		} else {
			r++
			count--
		}
		ans = MaxInts([]int{ans, count})
	}

	return ans
}
