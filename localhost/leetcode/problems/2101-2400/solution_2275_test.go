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

func TestSolution2275(t *testing.T) {
	assert := assert.New(t)
	actual := largestCombination([]int{16, 17, 71, 62, 12, 24, 14})
	assert.Equal(actual, 4, "SolutionOther2275")
}
