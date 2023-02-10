package solutions

import "unicode"

func rec_784(i int, s string, local_res []rune, res *[]string) {
	if i == len(s) {
		*res = append(*res, string(local_res))
		return
	}

	if unicode.IsLetter(rune(s[i])) {
		local_res[i] = unicode.ToLower(rune(s[i]))
		rec_784(i+1, s, local_res, res)
		local_res[i] = unicode.ToUpper(rune(s[i]))
		rec_784(i+1, s, local_res, res)
	} else {
		local_res[i] = rune(s[i])
		rec_784(i+1, s, local_res, res)
	}
}
func letterCasePermutation(s string) []string {
	n := len(s)
	local_res := make([]rune, n, n)
	res := make([]string, 0)

	rec_784(0, s, local_res, &res)

	return res
}
