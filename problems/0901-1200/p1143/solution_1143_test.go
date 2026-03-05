package p1143

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

func TestSolution1143(t *testing.T) {
	assert := assert.New(t)
	actual := longestCommonSubsequence("abcde", "ace")
	assert.Equal(actual, 3, "Solution1143")
}
