package p2109

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

func TestSolution2109(t *testing.T) {
	assert := assert.New(t)
	actual := addSpaces("LeetcodeHelpsMeLearn", []int{8, 13, 15})
	assert.Equal(actual, "Leetcode Helps Me Learn", "Solution2109")
}
