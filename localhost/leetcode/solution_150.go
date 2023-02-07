package solutions

import "strconv"

func evalRPN(tokens []string) int {
	stack := make([]int, 0)
	for _, token := range tokens {
		if token == "+" {
			operand1 := stack[len(stack)-1]
			operand2 := stack[len(stack)-2]
			stack = stack[:len(stack)-2]
			res := operand2 + operand1
			stack = append(stack, res)
		} else if token == "-" {
			operand1 := stack[len(stack)-1]
			operand2 := stack[len(stack)-2]
			stack = stack[:len(stack)-2]
			res := operand2 - operand1
			stack = append(stack, res)
		} else if token == "*" {
			operand1 := stack[len(stack)-1]
			operand2 := stack[len(stack)-2]
			stack = stack[:len(stack)-2]
			res := operand2 * operand1
			stack = append(stack, res)
		} else if token == "/" {
			operand1 := stack[len(stack)-1]
			operand2 := stack[len(stack)-2]
			stack = stack[:len(stack)-2]
			res := operand2 / operand1
			stack = append(stack, res)
		} else {
			num, _ := strconv.Atoi(token)
			stack = append(stack, num)
		}
	}

	return stack[0]
}
