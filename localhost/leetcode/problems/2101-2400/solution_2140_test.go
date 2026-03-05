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

func TestSolution2140(t *testing.T) {
	assert := assert.New(t)
	actual := mostPoints([][]int{{3, 2}, {4, 3}, {4, 4}, {2, 5}})
	assert.Equal(actual, int64(5), "Solution2140")
}
