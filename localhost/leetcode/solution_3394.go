package solutions

import "sort"

func checkValidCuts(n int, rectangles [][]int) bool {
	arrVertical := make([][]int, len(rectangles))
	arrHorizontal := make([][]int, len(rectangles))
	for i, r := range rectangles {
		arrVertical[i] = r
		arrHorizontal[i] = r
	}

	sort.Slice(arrVertical, func(i, j int) bool {
		return arrVertical[i][0] < arrVertical[j][0]
	})

	sort.Slice(arrHorizontal, func(i, j int) bool {
		return arrHorizontal[i][1] < arrHorizontal[j][1]
	})

	endX := arrVertical[0][2]
	resVert := 0
	for i := 1; i < len(arrVertical); i++ {
		startX := arrVertical[i][0]
		if startX >= endX {
			resVert++
		}

		endX = MaxInts([]int{endX, arrVertical[i][2]})
	}

	endY := arrHorizontal[0][3]
	resHor := 0
	for i := 1; i < len(arrHorizontal); i++ {
		startY := arrHorizontal[i][1]
		if startY >= endY {
			resHor++
		}
		endY = MaxInts([]int{endY, arrHorizontal[i][3]})
	}

	return resVert >= 2 || resHor >= 2
}
