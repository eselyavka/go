package solutions

import "math"

func longestSquareStreak(nums []int) int {
	set := make(map[int]struct{})
	for _, num := range nums {
		set[num] = struct{}{}
	}

	res := -1

	for _, num := range nums {
		p := 2
		buf := 1
		for true {
			if _, exists := set[int(math.Pow(float64(num), float64(p)))]; exists {
				buf++
				num = int(math.Pow(float64(num), float64(p)))
			} else {
				break
			}
		}

		res = MaxInts([]int{res, buf})
	}

	if res >= 2 {
		return res
	}

	return -1
}
