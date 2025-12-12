package solutions

func minOperations_3542(nums []int) int {
	stack := make([]int, 0)
	ans := 0

	for _, num := range nums {
		if num == 0 {
			stack = make([]int, 0)
			continue
		}

		for len(stack) > 0 && stack[len(stack)-1] > num {
			stack = stack[:len(stack)-1]
		}
		if len(stack) == 0 || stack[len(stack)-1] < num {
			ans++
			stack = append(stack, num)
		}
	}

	return ans
}
