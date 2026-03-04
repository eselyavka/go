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

func TestSolution1637(t *testing.T) {
	assert := assert.New(t)
	actual := maxWidthOfVerticalArea([][]int{{8, 7}, {9, 9}, {7, 4}, {9, 7}})
	assert.Equal(actual, 1, "Solution1637")
}
