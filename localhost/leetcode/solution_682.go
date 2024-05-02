package solutions

import "strconv"

func calPoints(operations []string) int {
	var stack []int

	for _, op := range operations {
		if op == "C" {
			stack = stack[:len(stack)-1]
		} else if op == "D" {
			stack = append(stack, stack[len(stack)-1]*2)
		} else if op == "+" {
			stack = append(stack, (stack[len(stack)-1] + stack[len(stack)-2]))
		} else {
			num, _ := strconv.Atoi(op)
			stack = append(stack, num)
		}
	}

	res := sum(stack)

	return res
}
