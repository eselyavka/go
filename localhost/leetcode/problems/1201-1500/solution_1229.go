package solutions

import "sort"

func minAvailableDuration(slots1 [][]int, slots2 [][]int, duration int) []int {
	sort.Slice(slots1, func(i, j int) bool {
		if len(slots1[i]) == 0 && len(slots1[j]) == 0 {
			return false
		}
		if len(slots1[i]) == 0 || len(slots1[j]) == 0 {
			return len(slots1[i]) == 0
		}

		return slots1[i][0] < slots1[j][0]
	})

	sort.Slice(slots2, func(i, j int) bool {
		if len(slots2[i]) == 0 && len(slots2[j]) == 0 {
			return false
		}
		if len(slots2[i]) == 0 || len(slots2[j]) == 0 {
			return len(slots2[i]) == 0
		}

		return slots2[i][0] < slots2[j][0]
	})

	s1 := len(slots1)
	s2 := len(slots2)

	p1 := 0
	p2 := 0

	for p1 < s1 && p2 < s2 {
		start1 := slots1[p1][0]
		end1 := slots1[p1][1]
		start2 := slots2[p2][0]
		end2 := slots2[p2][1]
		min := MinInts([]int{end1, end2})
		max := MaxInts([]int{start1, start2})
		diff := min - max
		if diff >= duration {
			return []int{max, max + duration}
		}

		if end1 < end2 {
			p1++
		} else {
			p2++
		}
	}

	return []int{}
}
