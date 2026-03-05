package solutions

import "sort"

func canAttendMeetings(intervals [][]int) bool {
	n := len(intervals)

	if n == 0 || n == 1 {
		return true
	}

	sort.Slice(intervals, func(i, j int) bool {
		if len(intervals[i]) == 0 && len(intervals[j]) == 0 {
			return false
		}
		if len(intervals[i]) == 0 || len(intervals[j]) == 0 {
			return len(intervals[i]) == 0
		}

		return intervals[i][0] < intervals[j][0]
	})

	prev_end := intervals[0][1]

	for i := 1; i < n; i++ {
		curr_start := intervals[i][0]
		curr_end := intervals[i][1]
		if prev_end > curr_start {
			return false
		}
		prev_end = curr_end
	}

	return true
}
