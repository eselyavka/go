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

func TestSolution1582(t *testing.T) {
	assert := assert.New(t)
	actual := numSpecial([][]int{{1, 0, 0}, {0, 0, 1}, {1, 0, 0}})
	assert.Equal(actual, 1, "SolutionOther1582")
}
