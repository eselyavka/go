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

func TestSolution2055(t *testing.T) {
	assert := assert.New(t)
	actual := platesBetweenCandles("**|**|***|", [][]int{{2, 5}, {5, 9}})
	assert.Equal([]int{2, 3}, actual, "Solution2055")
}
