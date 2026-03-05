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

func TestSolution2475(t *testing.T) {
	assert := assert.New(t)
	actual := unequalTriplets([]int{4, 4, 2, 4, 3})
	assert.Equal(3, actual, "Solution2475")
}
