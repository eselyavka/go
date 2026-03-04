package solutions

import "math"

func areaOfMaxDiagonal(dimensions [][]int) int {
	diagLen := -1.0
	ans := -1

	for _, d := range dimensions {
		length := d[0]
		width := d[1]
		currDiagLen := math.Sqrt(float64(length*length + width*width))
		if currDiagLen > diagLen {
			diagLen = currDiagLen
			ans = length * width
		} else if currDiagLen == diagLen && ans < length*width {
			ans = length * width
		}
	}

	return ans
}
