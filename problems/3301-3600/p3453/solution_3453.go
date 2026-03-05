package p3453

import "math"

func separateSquares(squares [][]int) float64 {
	totalArea := 0.0
	low := math.Inf(1)
	high := math.Inf(-1)

	for _, sq := range squares {
		y := float64(sq[1])
		l := float64(sq[2])

		totalArea += l * l
		if y < low {
			low = y
		}
		if y+l > high {
			high = y + l
		}
	}

	target := totalArea / 2.0

	var h float64
	for iter := 0; iter < 60; iter++ {
		h = (low + high) / 2.0
		below := 0.0

		for _, sq := range squares {
			y := float64(sq[1])
			l := float64(sq[2])

			if h <= y {
				continue
			} else if h >= y+l {
				below += l * l
			} else {
				below += l * (h - y)
			}
		}

		if below >= target {
			high = h
		} else {
			low = h
		}
	}

	return h
}
