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

func TestSolution910(t *testing.T) {
	assert := assert.New(t)
	actual := smallestRangeII([]int{1, 3, 6}, 3)
	assert.Equal(3, actual, "Solution910")
}
