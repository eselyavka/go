package solutions

func largestGoodInteger(num string) string {
	n := len(num)

	res := ""

	for i := 0; i < n-2; i++ {
		s := num[i : i+3]
		if (s[0] == s[1]) && (s[0] == s[2]) && s > res {
			res = s
		}
	}

	return res
}
