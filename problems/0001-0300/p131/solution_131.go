package p131

import "github.com/eseliavka/go/util"

func backtracking(local_ans []string, pos int, s string, ans *[][]string) {
	if pos >= len(s) {
		tmp := make([]string, len(local_ans))
		copy(tmp, local_ans)
		*ans = append(*ans, tmp)
		return
	}

	for i := pos; i < len(s); i++ {
		if util.IsPalindrome(s, pos, i) {
			local_ans = append(local_ans, s[pos:i+1])
			backtracking(local_ans, i+1, s, ans)
			local_ans = local_ans[:len(local_ans)-1]
		}
	}
}

func partition(s string) [][]string {
	ans := make([][]string, 0)

	backtracking([]string{}, 0, s, &ans)

	return ans
}
