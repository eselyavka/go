package p70

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

func TestSolution70(t *testing.T) {
	assert := assert.New(t)
	actual := climbStairs(4)
	assert.Equal(actual, 5, "Solution70")
}
