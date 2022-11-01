package solutions

func maxArea(height []int) int {
	max_square := -1
	l := 0
	r := len(height) - 1
	square := -1

	for l <= r {
		if height[r] >= height[l] {
			square = height[l] * (r - l)
			if max_square <= square {
				max_square = square
			}
			l++
		} else {
			square = height[r] * (r - l)
			if max_square <= square {
				max_square = square
			}
			r--
		}
	}

	return max_square
}
