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

func TestSolution2490(t *testing.T) {
	assert := assert.New(t)
	actual := isCircularSentence("leetcode eats soul")
	assert.True(actual, "Solution2490")
}
