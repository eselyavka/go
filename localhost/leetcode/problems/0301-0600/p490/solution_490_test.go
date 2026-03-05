package p490

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

func TestSolution490(t *testing.T) {
	assert := assert.New(t)
	actual := hasPath([][]int{{0, 0, 1, 0, 0}, {0, 0, 0, 0, 0}, {0, 0, 0, 1, 0}, {1, 1, 0, 1, 1}, {0, 0, 0, 0, 0}}, []int{0, 4}, []int{4, 4})
	assert.True(actual, "Solution490")
}
