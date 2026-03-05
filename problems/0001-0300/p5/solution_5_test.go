package p5

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

func TestSolution5(t *testing.T) {
	assert := assert.New(t)
	res := longestPalindrome("babad")
	assert.Equal(res, "bab", "Solution5")

	res = longestPalindromeDP("babad")
	assert.Equal(res, "bab", "Solution5")
}
