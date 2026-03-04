package solutions

import (
	"sort"
	"strings"
)

func groupAnagrams(strs []string) [][]string {
	var res [][]string
	var ans = make(map[string][]string)
	for _, str := range strs {
		chars := strings.Split(str, "")
		sort.Strings(chars)
		sorted_string := strings.Join(chars, "")
		val, ok := ans[sorted_string]
		if ok {
			ans[sorted_string] = append(val, str)
		} else {
			ans[sorted_string] = []string{str}
		}
	}
	for _, value := range ans {
		res = append(res, value)
	}
	return res
}
