package solutions

import "sort"

func findMinArrowShots(points [][]int) int {
	n := len(points)

	if n == 1 {
		return 1
	}

	sort.Slice(points, func(i, j int) bool {
		if len(points[i]) == 0 && len(points[j]) == 0 {
			return false
		}
		if len(points[i]) == 0 || len(points[j]) == 0 {
			return len(points[i]) == 0
		}

		return points[i][1] < points[j][1]
	})

	ans := 1

	pxe := points[0][1]
	for i := 1; i < n; i++ {
		xs := points[i][0]
		xe := points[i][1]

		if xs <= pxe {
			continue
		}
		pxe = xe
		ans++
	}

	return ans
}
