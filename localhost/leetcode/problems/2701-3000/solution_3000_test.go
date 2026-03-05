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

func TestSolution3000(t *testing.T) {
	assert := assert.New(t)
	actual := areaOfMaxDiagonal([][]int{{9, 3}, {8, 6}})
	assert.Equal(actual, 48, "Solution3000")
}
