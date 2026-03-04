package solutions

func minimumIndex(nums []int) int {
	n := len(nums)

	counter := make(map[int]int)
	for _, num := range nums {
		counter[num]++
	}

	maxOccur := -1
	candidate := -1
	for k, v := range counter {
		if v > maxOccur {
			maxOccur = v
			candidate = k
		}
	}

	leftOccur := 0
	remaining := maxOccur
	for i, val := range nums {
		if val == candidate {
			remaining--
			leftOccur++

			if leftOccur > (i+1-leftOccur) && remaining > ((n-(i+1))-remaining) {
				return i
			}
		}
	}

	return -1
}
