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

func TestSolution2501(t *testing.T) {
	assert := assert.New(t)
	actual := longestSquareStreak([]int{4, 3, 6, 16, 8, 2})
	assert.Equal(actual, 3, "Solution2501")
}
