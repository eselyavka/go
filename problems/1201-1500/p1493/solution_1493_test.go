package p1493

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

func TestSolution1493(t *testing.T) {
	assert := assert.New(t)
	actual := longestSubarray([]int{1, 1, 0, 1})
	assert.Equal(actual, 3, "Solution1493")
}
