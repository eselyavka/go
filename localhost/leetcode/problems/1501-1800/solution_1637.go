package solutions

import "sort"

func maxWidthOfVerticalArea(points [][]int) int {
	xArr := make([]int, len(points))

	for pos, point := range points {
		x := point[0]
		xArr[pos] = x
	}

	sort.Ints(xArr)

	ans := 0

	for i := 1; i < len(xArr); i++ {
		ans = MaxInts([]int{ans, xArr[i] - xArr[i-1]})
	}

	return ans
}
