package p2845

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

func TestSolution2845(t *testing.T) {
	assert := assert.New(t)
	actual := countInterestingSubarrays([]int{3, 2, 4}, 2, 1)
	assert.Equal(actual, int64(3), "Solution2845")
}
