package solutions

func isValid(s string) bool {
	mapping := map[string]string{
		")": "(",
		"}": "{",
		"]": "[",
	}
	stack := make([]string, 0)

	for _, val := range s {
		c := string(val)
		if c == "(" || c == "{" || c == "[" {
			stack = append(stack, c)
		} else {
			if len(stack) == 0 {
				return false
			}

			o := stack[len(stack)-1]
			stack = stack[:len(stack)-1]

			if val, ok := mapping[c]; ok {
				if val != o {
					return false
				}
			} else {
				return false
			}
		}
	}

	if len(stack) > 0 {
		return false
	}

	return true
}
