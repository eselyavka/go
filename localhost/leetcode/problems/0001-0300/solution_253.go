package solutions

import "sort"

func minMeetingRooms(intervals [][]int) int {
	n := len(intervals)
	start_points := make([]int, n, n)
	end_points := make([]int, n, n)

	i := 0
	for _, interval := range intervals {
		start_points[i] = interval[0]
		end_points[i] = interval[1]
		i++
	}

	sort.Ints(start_points)
	sort.Ints(end_points)

	s := 0
	e := 0
	ans := 0
	count := 0

	for s < n {
		if start_points[s] < end_points[e] {
			count++
			s++
		} else {
			count--
			e++
		}
		ans = MaxInts([]int{ans, count})
	}

	return ans
}
