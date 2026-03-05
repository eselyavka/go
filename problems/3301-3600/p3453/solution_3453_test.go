package p3453

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

func TestSolution3453(t *testing.T) {
	assert := assert.New(t)
	actual := separateSquares([][]int{{0, 0, 1}, {2, 2, 1}})
	assert.Equal(actual, 1.0, "SolutionOther3453")
}
