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

func TestSolution3446(t *testing.T) {
	assert := assert.New(t)
	actual := sortMatrix([][]int{{1, 7, 3}, {9, 8, 2}, {4, 5, 6}})
	assert.Equal(actual, [][]int{{8, 2, 3}, {9, 6, 7}, {4, 5, 1}}, "Solution3446")
}
