package solutions

import (
	"bytes"
)

func getHash(str string, shift int32) string {
	var buf bytes.Buffer
	for _, c := range str {
		code := (c - shift)
		if code < 0 {
			code = code + 26
		}

		buf.WriteString(string(code + []rune("a")[0]))
	}

	return buf.String()
}

func groupStrings(strings []string) [][]string {
	acc := make(map[string][]string)

	for _, str := range strings {
		hash := getHash(str, []rune(str[0:1])[0])
		acc[hash] = append(acc[hash], str)
	}

	res := make([][]string, 0)

	for _, v := range acc {
		res = append(res, v)
	}

	return res
}
