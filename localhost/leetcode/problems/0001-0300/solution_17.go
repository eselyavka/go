package solutions

func backtrack_17(idx int, acc, digits string, mapping map[uint8][]string, ans *[]string) {
	if len(acc) == len(digits) {
		*ans = append(*ans, acc)
		return
	}
	for _, c := range mapping[digits[idx]] {
		acc += c
		backtrack_17(idx+1, acc, digits, mapping, ans)
		acc = acc[:len(acc)-1]
	}
}
func letterCombinations(digits string) []string {
	if digits == "" {
		return []string{}
	}

	mapping := make(map[uint8][]string)
	mapping['2'] = []string{"a", "b", "c"}
	mapping['3'] = []string{"d", "e", "f"}
	mapping['4'] = []string{"g", "h", "i"}
	mapping['5'] = []string{"j", "k", "l"}
	mapping['6'] = []string{"m", "n", "o"}
	mapping['7'] = []string{"p", "q", "r", "s"}
	mapping['8'] = []string{"t", "u", "v"}
	mapping['9'] = []string{"w", "x", "y", "z"}

	ans := make([]string, 0)
	local_res := ""

	backtrack_17(0, local_res, digits, mapping, &ans)

	return ans
}
