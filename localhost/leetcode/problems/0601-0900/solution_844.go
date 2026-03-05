package solutions

import "strings"

func backspaceCompare(s string, t string) bool {
	var stack_s []string
	for _, c := range s {
		if c == '#' {
			if len(stack_s) >= 1 {
				stack_s = stack_s[:len(stack_s)-1]
			}

		} else {
			stack_s = append(stack_s, string(c))
		}

	}

	var stack_t []string
	for _, c := range t {
		if c == '#' {
			if len(stack_t) >= 1 {
				stack_t = stack_t[:len(stack_t)-1]
			}
		} else {
			stack_t = append(stack_t, string(c))
		}

	}

	return strings.Join(stack_s, "") == strings.Join(stack_t, "")
}
