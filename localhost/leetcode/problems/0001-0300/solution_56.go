package solutions

import "sort"

func mergeIntervals(intervals [][]int) [][]int {
	n := len(intervals)

	if n == 1 {
		return intervals
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

	prev_s := intervals[0][0]
	prev_e := intervals[0][1]

	for i := 1; i < n; i++ {
		curr_s := intervals[i][0]
		curr_e := intervals[i][1]
		if curr_s <= prev_e {
			new_int := []int{MinInts([]int{curr_s, prev_s}), MaxInts([]int{curr_e, prev_e})}
			intervals[i] = new_int
			intervals[i-1] = []int{}
			prev_s = new_int[0]
			prev_e = new_int[1]
			continue
		}
		prev_s = curr_s
		prev_e = curr_e
	}

	ans := make([][]int, 0)
	for _, interval := range intervals {
		if len(interval) != 0 {
			ans = append(ans, interval)
		}
	}
	return ans
}
