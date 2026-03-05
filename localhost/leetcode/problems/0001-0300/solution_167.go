package solutions

func twoSum_167(numbers []int, target int) []int {
	n := len(numbers)

	for i := 0; i < n; i++ {
		diff := target - numbers[i]
		idx := binarySearch(0, n-1, numbers, diff)
		if idx >= 0 && idx != i {
			return []int{MinInts([]int{i, idx}) + 1, MaxInts([]int{i, idx}) + 1}
		}
	}

	return []int{}
}
