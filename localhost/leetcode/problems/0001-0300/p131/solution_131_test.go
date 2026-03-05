package p131

import (
	"github.com/stretchr/testify/assert"
	"localhost/leetcode/util"
	"sort"
	"strings"
	"testing"
)

var (
	_ = sort.Strings
	_ = strings.Join
)

func TestSolution131(t *testing.T) {
	assert := assert.New(t)
	actual := partition("aab")
	expected := [][]string{{"a", "a", "b"}, {"aa", "b"}}
	assert.Equal(len(actual), len(expected), "Solution131")
	for _, left := range expected {
		is_equal := false
		for _, right := range actual {
			if util.StringSlicesEqual(left, right) {
				is_equal = true
			}
		}
		assert.True(is_equal, "Solution131")
	}
}
