package solutions

import (
	"sort"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

var (
	_ = sort.Strings
	_ = strings.Join
)

func TestSolution49(t *testing.T) {
	assert := assert.New(t)
	tests := []struct {
		expected []string
	}{
		{[]string{"bat"}},
		{[]string{"nat", "tan"}},
		{[]string{"ate", "eat", "tea"}},
	}

	var arr = []string{"eat", "tea", "tan", "ate", "nat", "bat"}
	actual := groupAnagrams(arr)
	for _, expected := range tests {
		is_equal := false
		for _, actual := range actual {
			sort.Strings(expected.expected)
			sort.Strings(actual)
			if strings.Join(expected.expected, "") == strings.Join(actual, "") {
				is_equal = true
			}
		}
		assert.True(is_equal, "Solution49")
	}
}
