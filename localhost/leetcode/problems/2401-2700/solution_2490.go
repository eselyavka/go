package solutions

import "strings"

func isCircularSentence(sentence string) bool {
	tokens := strings.Fields(sentence)

	if len(tokens) == 1 {
		return tokens[0][0] == tokens[0][len(tokens[0])-1]
	}

	if tokens[0][0] != tokens[len(tokens)-1][len(tokens[len(tokens)-1])-1] {
		return false
	}

	s := tokens[0][len(tokens[0])-1]
	for i := 1; i < len(tokens); i++ {
		e := tokens[i][0]
		if s != e {
			return false
		}
		s = tokens[i][len(tokens[i])-1]
	}

	return true
}
