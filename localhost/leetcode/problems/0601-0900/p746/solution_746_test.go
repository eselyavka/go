package p746

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

func TestSolution746(t *testing.T) {
	assert := assert.New(t)
	actual := minCostClimbingStairs([]int{1, 100, 1, 1, 1, 100, 1, 1, 100, 1})
	assert.Equal(actual, 6, "Solution746")
}
