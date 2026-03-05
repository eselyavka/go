package solutions

import "strings"

func isPrefixOfWord(sentence string, searchWord string) int {
	tokens := strings.Split(sentence, " ")
	for idx, token := range tokens {
		if strings.HasPrefix(token, searchWord) {
			return idx + 1
		}
	}
	return -1
}
