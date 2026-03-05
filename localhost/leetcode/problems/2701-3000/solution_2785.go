package solutions

import (
	"sort"
	"strings"
)

func sortVowels(s string) string {
	n := len(s)

	vowels := []string{"a", "e", "i", "o", "u"}

	ans := make([]string, n, n)

	vowels_acc := make([]rune, 0)

	for i, c := range s {
		if !inStringArray(vowels, strings.ToLower(string(c))) {
			ans[i] = string(c)
		} else {
			vowels_acc = append(vowels_acc, c)
		}
	}

	sort.Slice(vowels_acc, func(i, j int) bool { return vowels_acc[i] < vowels_acc[j] })

	j := 0
	for i := 0; i < n; i++ {
		if ans[i] == "" {
			ans[i] = string(vowels_acc[j])
			j++
		}
	}

	return strings.Join(ans, "")
}
