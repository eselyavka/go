package solutions

import "strings"

func rec_22(curr_idx, m, balance int, buf []string, ans *[]string) {
	if curr_idx == m {
		if balance == 0 {
			*ans = append(*ans, strings.Join(buf, ""))
		}
		return
	}

	if balance > 0 {
		buf[curr_idx] = ")"
		rec_22(curr_idx+1, m, balance-1, buf, ans)
	}

	buf[curr_idx] = "("
	rec_22(curr_idx+1, m, balance+1, buf, ans)
}

func generateParenthesis(n int) []string {
	m := 2 * n
	ans := make([]string, 0)
	buf := make([]string, m, m)

	rec_22(0, m, 0, buf, &ans)

	return ans
}
