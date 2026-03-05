package p159

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

func TestSolution159(t *testing.T) {
	assert := assert.New(t)
	actual := lengthOfLongestSubstringTwoDistinct("ccaabbb")
	assert.Equal(5, actual, "Solution159")
}
