package p2452

func withinTwo(s1, s2 string) bool {
	edits := 0
	for i := 0; i < len(s1); i++ {
		if s1[i] != s2[i] {
			edits++
			if edits > 2 {
				return false
			}
		}
	}

	return true
}

func twoEditWords(queries []string, dictionary []string) []string {
	ans := make([]string, 0)
	for _, query := range queries {
		for _, word := range dictionary {
			if withinTwo(query, word) {
				ans = append(ans, query)
				break
			}
		}
	}
	return ans
}
