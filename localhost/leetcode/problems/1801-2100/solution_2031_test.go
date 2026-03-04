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

func TestSolution2031(t *testing.T) {
	assert := assert.New(t)
	actual := subarraysWithMoreZerosThanOnes([]int{0, 1, 1, 0, 1})
	assert.Equal(actual, 9, "Solution2031")
}
