package p494

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

func TestSolution494(t *testing.T) {
	assert := assert.New(t)
	actual := findTargetSumWays([]int{1, 1, 1, 1, 1}, 3)
	assert.Equal(actual, 5, "Solution494")
}
