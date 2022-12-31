package solutions

import "sort"

func eraseOverlapIntervals(intervals [][]int) int {
	n := len(intervals)
	if n == 0 {
		return 0
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

	prev_e := intervals[0][1]
	ans := 0
	for i := 1; i < n; i++ {
		curr_s := intervals[i][0]
		curr_e := intervals[i][1]
		if prev_e > curr_s {
			prev_e = MinInts([]int{prev_e, curr_e})
			ans++
		} else {
			prev_e = curr_e
		}
	}
	return ans
}
