package p1536

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

func TestSolution1536(t *testing.T) {
	assert := assert.New(t)
	actual := minSwaps([][]int{{0, 0, 1}, {1, 1, 0}, {1, 0, 0}})
	assert.Equal(actual, 3, "SolutionOther1536")
}
