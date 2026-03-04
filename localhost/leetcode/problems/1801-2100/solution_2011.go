package solutions

import "strings"

func finalValueAfterOperations(operations []string) int {

	ans := 0
	for _, operation := range operations {
		if strings.Contains(operation, "++") {
			ans++
		} else {
			ans--
		}
	}

	return ans
}
