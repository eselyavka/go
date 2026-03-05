package p1930

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

func TestSolution1930(t *testing.T) {
	assert := assert.New(t)
	actual := countPalindromicSubsequence("bbcbaba")
	assert.Equal(actual, 4, "Solution1930")
}
