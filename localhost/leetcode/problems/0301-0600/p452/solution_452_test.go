package p452

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

func TestSolution452(t *testing.T) {
	assert := assert.New(t)
	actual := findMinArrowShots([][]int{{10, 16}, {2, 8}, {1, 6}, {7, 12}})
	assert.Equal(2, actual, "Solution452")
}
