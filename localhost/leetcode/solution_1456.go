package solutions

func maxVowels(s string, k int) int {
	vowel := map[rune]struct{}{
		'a': {},
		'e': {},
		'i': {},
		'o': {},
		'u': {},
	}

	cnt := 0

	for i := 0; i < k; i++ {
		_, ok := vowel[rune(s[i])]
		if ok {
			cnt++
		}
	}

	ans := cnt
	n := len(s)

	for i := k; i < n; i++ {
		_, ok := vowel[rune(s[i])]
		if ok {
			cnt++
		}

		_, ok = vowel[rune(s[i-k])]

		if ok {
			cnt--
		}

		ans = MaxInts([]int{ans, cnt})
	}

	return ans
}
