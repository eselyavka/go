package solutions

import "sort"

func countDays(days int, meetings [][]int) int {
	sort.Slice(meetings, func(i, j int) bool { return meetings[i][0] < meetings[j][0] })
	mergedMeetings := make([][]int, 0)
	start := meetings[0][0]
	end := meetings[0][1]

	for i := 1; i < len(meetings); i++ {
		currStart := meetings[i][0]
		currEnd := meetings[i][1]

		if currStart <= end {
			end = MaxInts([]int{end, currEnd})
		} else {
			mergedMeetings = append(mergedMeetings, []int{start, end})
			start = currStart
			end = currEnd
		}
	}

	mergedMeetings = append(mergedMeetings, []int{start, end})
	ans := 0

	ans += MaxInts([]int{mergedMeetings[0][0] - 1, 0})

	for i := 1; i < len(mergedMeetings); i++ {
		prevEnd := mergedMeetings[i-1][1]
		currStart := mergedMeetings[i][0]
		ans += MaxInts([]int{currStart - prevEnd - 1, 0})
	}
	ans += MaxInts([]int{days - mergedMeetings[len(mergedMeetings)-1][1], 0})

	return ans
}
