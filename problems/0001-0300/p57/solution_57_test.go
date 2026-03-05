package p57

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

func TestSolution57(t *testing.T) {
	assert := assert.New(t)
	actual := insert([][]int{{1, 2}, {3, 5}, {6, 7}, {8, 10}, {12, 16}}, []int{4, 8})
	assert.Equal([][]int{{1, 2}, {3, 10}, {12, 16}}, actual, "Solution57")
}
