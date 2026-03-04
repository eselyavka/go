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

func TestSolution54(t *testing.T) {
	assert := assert.New(t)
	actual := spiralOrder([][]int{{1, 2, 3}, {4, 5, 6}, {7, 8, 9}})
	assert.Equal([]int{1, 2, 3, 6, 9, 8, 7, 4, 5}, actual, "Solution54")
}
