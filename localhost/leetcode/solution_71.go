package solutions

import "strings"

func simplifyPath(path string) string {
	stack := make([]string, 0)
	tokens := strings.Split(path, "/")

	ans := "/"

	for _, token := range tokens {
		if token == "" {
			continue
		} else if token == "." {
			continue
		} else if token == ".." {
			if len(stack) > 0 {
				stack = stack[:len(stack)-1]
			}
		} else {
			stack = append(stack, token)
		}
	}

	ans += strings.Join(stack, "/")

	return ans
}
