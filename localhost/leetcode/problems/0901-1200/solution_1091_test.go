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

func TestSolution1091(t *testing.T) {
	assert := assert.New(t)
	res := shortestPathBinaryMatrix([][]int{{0, 1}, {1, 0}})
	assert.Equal(res, 2, "Solution1091")
}
