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

func TestSolution2654(t *testing.T) {
	assert := assert.New(t)
	actual := minOperations_2654([]int{2, 6, 3, 4})
	assert.Equal(actual, 4, "SolutionOther2654")
}
