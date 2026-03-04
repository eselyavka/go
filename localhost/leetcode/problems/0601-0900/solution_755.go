package solutions

func pourWater(heights []int, volume int, k int) []int {
	n := len(heights)

	for volume > 0 {
		idx := k

		for idx > 0 && heights[idx] >= heights[idx-1] {
			idx--
		}

		for idx < n-1 && heights[idx] >= heights[idx+1] {
			idx++
		}

		for idx > k && heights[idx] >= heights[idx-1] {
			idx--
		}

		heights[idx]++
		volume -= 1
	}

	return heights
}
