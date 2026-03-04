package solutions

func maxArea(height []int) int {
	max_square := -1
	l := 0
	r := len(height) - 1

	for l <= r {
		if height[r] >= height[l] {
			curr := height[l] * (r - l)
			if max_square <= curr {
				max_square = curr
			}
			l++
		} else {
			curr := height[r] * (r - l)
			if max_square <= curr {
				max_square = curr
			}
			r--
		}
	}

	return max_square
}
