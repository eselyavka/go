package solutions

import "unicode"

func calculate(s string) int {
	digit_stack := make([]int, 0)
	sign := "+"
	num := 0

	for i, c := range s {
		if unicode.IsDigit(c) {
			num = num*10 + int(c-'0')
		}

		if i+1 == len(s) || (string(c) == "+" || string(c) == "-" || string(c) == "*" || string(c) == "/") {
			if sign == "+" {
				digit_stack = append(digit_stack, num)
			} else if sign == "-" {
				digit_stack = append(digit_stack, -1*num)
			} else if sign == "*" {
				digit_stack[len(digit_stack)-1] = digit_stack[len(digit_stack)-1] * num
			} else {
				digit_stack[len(digit_stack)-1] = digit_stack[len(digit_stack)-1] / num
			}
			num = 0
			sign = string(c)
		}
	}

	return sum(digit_stack)
}
