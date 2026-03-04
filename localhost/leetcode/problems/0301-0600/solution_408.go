package solutions

import (
	"strconv"
	"unicode"
)

func validWordAbbreviation(word string, abbr string) bool {
	n := len(word)
	m := len(abbr)

	if strconv.FormatInt(int64(n), 10) == abbr {
		return true
	}

	i := 0
	j := 0

	for i < n && j < m {
		if word[i] == abbr[j] {
			i += 1
			j += 1
			continue
		}

		if (word[i] != abbr[j]) && unicode.IsLetter(rune(abbr[j])) {
			return false
		}

		if word[i] != abbr[j] {
			num := 0
			for j < m && unicode.IsDigit(rune(abbr[j])) {
				if num == 0 && abbr[j] == 0x30 {
					return false
				}
				buf, _ := strconv.Atoi(string(abbr[j]))
				num = num*10 + buf
				j += 1
			}
			i += num
		}
	}

	return i == n && j == m
}
