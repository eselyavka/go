package solutions

import "math"

func maxBottlesDrunk(numBottles int, numExchange int) int {
	E0 := numBottles
	x := numExchange

	if E0 <= 0 {
		return 0
	}

	D := math.Pow(float64(2*x-3), 2.0) + float64(8*(E0-1))

	s := math.Sqrt(D)
	t := math.Floor((s - float64(2*x-3)) / float64(2))

	if t < 0 {
		t = 0
	}

	return numBottles + int(t)
}
