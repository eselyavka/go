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

func TestSolution64(t *testing.T) {
	assert := assert.New(t)
	actual := minPathSum([][]int{{1, 3, 1}, {1, 5, 1}, {4, 2, 1}})
	assert.Equal(7, actual, "Solution64")
}
