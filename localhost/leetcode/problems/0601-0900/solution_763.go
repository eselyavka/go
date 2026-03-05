package solutions

import "sort"

func partitionLabels(s string) []int {
	last := make(map[rune][]int)

	for i, c := range s {
		last[c] = append(last[c], i)
	}

	var intervals [][]int
	for _, positions := range last {
		intervals = append(intervals, []int{positions[0], positions[len(positions)-1]})
	}

	sort.Slice(intervals, func(i, j int) bool {
		return intervals[i][0] < intervals[j][0]
	})

	mergedIntervals := [][]int{}
	start, end := intervals[0][0], intervals[0][1]

	for i := 1; i < len(intervals); i++ {
		currStart, currEnd := intervals[i][0], intervals[i][1]
		if currStart <= end {
			end = MaxInts([]int{end, currEnd})
		} else {
			mergedIntervals = append(mergedIntervals, []int{start, end})
			start, end = currStart, currEnd
		}
	}

	mergedIntervals = append(mergedIntervals, []int{start, end})

	var ans []int
	for _, t := range mergedIntervals {
		ans = append(ans, t[1]-t[0]+1)
	}

	return ans
}
