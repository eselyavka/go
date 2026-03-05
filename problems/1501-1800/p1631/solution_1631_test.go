package p1631

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

func TestSolution1631(t *testing.T) {
	assert := assert.New(t)

	actual := minimumEffortPath([][]int{{1, 2, 2}, {3, 8, 2}, {5, 3, 5}})

	assert.Equal(2, actual, "Solution1631")
}
