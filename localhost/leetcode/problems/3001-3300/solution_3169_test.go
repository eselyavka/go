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

func TestSolution3169(t *testing.T) {
	assert := assert.New(t)
	actual := countDays(10, [][]int{{5, 7}, {1, 3}, {9, 10}})
	assert.Equal(actual, 2, "Solution3169")
}
