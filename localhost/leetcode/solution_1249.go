package solutions

import "strings"

func minRemoveToMakeValid(s string) string {
	stack_o := make([]string, len(s))
	stack_c := make([]string, len(s))
	buf := make([]string, len(s))
	open_p := []byte("(")[0]
	close_p := []byte(")")[0]

	for i := 0; i < len(s); i++ {
		c := s[i]
		if c == open_p {
			stack_o = append(stack_o, string(c))
			buf = append(buf, string(c))
		} else if c == close_p {
			stack_c = append(stack_c, string(c))
			if len(stack_c) > len(stack_o) {
				stack_c = stack_c[:len(stack_c)-1]
				continue
			}
			buf = append(buf, string(c))
		} else {
			buf = append(buf, string(c))
		}
	}

	i := len(buf) - 1
	stack_o = make([]string, len(buf))
	stack_c = make([]string, len(buf))
	for i >= 0 {
		if buf[i] == ")" {
			stack_c = append(stack_c, buf[i])
		} else if buf[i] == "(" {
			stack_o = append(stack_o, buf[i])
			if len(stack_o) > len(stack_c) {
				stack_o = stack_o[:len(stack_o)-1]
				buf[i] = ""
			}
		}
		i -= 1
	}

	res := strings.Join(buf, "")
	return res
}
