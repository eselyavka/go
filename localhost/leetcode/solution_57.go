package solutions

func binary_search_57(intervals [][]int, newInterval []int) int {
	left := 0
	right := len(intervals) - 1
	var mid int
	for left <= right {
		mid = (left + right) / 2
		if intervals[mid][0] == newInterval[0] {
			return mid
		}

		if intervals[mid][0] < newInterval[0] {
			left = mid + 1
		} else {
			right = mid - 1
		}
	}

	return right + 1
}

func insert_57(intervals [][]int, newInterval []int) [][]int {
	if len(intervals) == 0 {
		return [][]int{newInterval}
	}

	idx := binary_search_57(intervals, newInterval)

	unmerged_intevals := make([][]int, 0)

	for _, val := range intervals[:idx] {
		unmerged_intevals = append(unmerged_intevals, val)
	}

	unmerged_intevals = append(unmerged_intevals, newInterval)

	for _, val := range intervals[idx:] {
		unmerged_intevals = append(unmerged_intevals, val)
	}

	n := len(unmerged_intevals)

	prev_start := unmerged_intevals[0][0]
	prev_end := unmerged_intevals[0][1]

	for i := 1; i < n; i++ {
		curr_start := unmerged_intevals[i][0]
		curr_end := unmerged_intevals[i][1]
		if curr_start <= prev_end {
			new_int := []int{MinInts([]int{prev_start, curr_start}), MaxInts([]int{prev_end, curr_end})}
			unmerged_intevals[i] = new_int
			unmerged_intevals[i-1] = []int{}
			prev_start = new_int[0]
			prev_end = new_int[1]
			continue
		}
		prev_start = curr_start
		prev_end = curr_end

	}

	ans := make([][]int, 0)
	n = len(unmerged_intevals)
	for i := 0; i < n; i++ {
		if len(unmerged_intevals[i]) != 0 {
			ans = append(ans, unmerged_intevals[i])
		}
	}
	return ans
}
