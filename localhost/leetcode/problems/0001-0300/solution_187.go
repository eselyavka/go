package solutions

func findRepeatedDnaSequences(s string) []string {
	n := len(s)

	if n <= 10 {
		return []string{}
	}

	l := 0
	r := 9
	hashmap := make(map[string]struct{})
	acc := make(map[string]struct{})

	for r < n {
		sub := s[l : r+1]
		if _, ok := hashmap[sub]; ok {
			acc[sub] = struct{}{}
		}
		hashmap[sub] = struct{}{}
		l++
		r++
	}
	ans := make([]string, 0)

	for key, _ := range acc {
		ans = append(ans, key)
	}

	return ans
}
