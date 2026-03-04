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

func TestSolution2384(t *testing.T) {
	assert := assert.New(t)
	actual := largestPalindromic("444947137")
	assert.Equal("7449447", actual, "Solution2384")
}
