package solutions

func vowelStrings(words []string, queries [][]int) []int {
	prefixSum := make([]int, len(words))
	vowels := make(map[rune]struct{})
	for _, vowel := range []rune{'a', 'e', 'i', 'o', 'u'} {
		vowels[vowel] = struct{}{}
	}
	vowelCounter := 0
	for idx, word := range words {
		if _, ok := vowels[rune(word[0])]; ok {
			if _, ok := vowels[rune(word[len(word)-1])]; ok {
				vowelCounter++
			}
		}
		prefixSum[idx] = vowelCounter
	}
	ans := make([]int, 0)
	for _, query := range queries {
		l := query[0]
		r := query[1]
		if l == 0 {
			ans = append(ans, prefixSum[r])
		} else {
			ans = append(ans, prefixSum[r]-prefixSum[l-1])
		}
	}
	return ans
}
