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

func TestSolution2110(t *testing.T) {
	assert := assert.New(t)
	actual := getDescentPeriods([]int{3, 2, 1, 4})
	assert.Equal(actual, int64(7), "SolutionOther2110")
}
