package p1

// twoSum returns indices of the two numbers such that they add up to target.
func twoSum(nums []int, target int) []int {
	result := make([]int, 0, 2)
	seenByValue := make(map[int]int, len(nums))

	for i, num := range nums {
		if j, ok := seenByValue[target-num]; ok {
			result = append(result, j, i)
		}
		seenByValue[num] = i
	}

	return result
}
