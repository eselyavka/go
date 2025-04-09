package solutions

func minOperations_3375(nums []int, k int) int {
	if len(nums) == 0 {
		return -1
	}

	minVal := MinInts(nums)

	if minVal < k {
		return -1
	}

	set := make(map[int]bool)
	for _, num := range nums {
		set[num] = true
	}

	if set[k] {
		return len(set) - 1
	}

	return len(set)
}
