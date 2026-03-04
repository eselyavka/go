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

func TestSolution3(t *testing.T) {
	assert := assert.New(t)
	res := lengthOfLongestSubstring("abcabcbb")
	assert.Equal(res, 3, "Solution3")
}
