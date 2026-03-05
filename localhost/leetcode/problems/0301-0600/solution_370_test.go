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

func TestSolution370(t *testing.T) {
	assert := assert.New(t)
	actual := getModifiedArray(5, [][]int{{1, 3, 2}, {2, 4, 3}, {0, 2, -2}})
	assert.Equal([]int{-2, 0, 3, 5, 3}, actual, "Solution370")
}
