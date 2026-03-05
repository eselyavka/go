package p1984

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

func TestSolution1984(t *testing.T) {
	assert := assert.New(t)
	actual := minimumDifference([]int{9, 4, 1, 7}, 2)
	assert.Equal(actual, 2, "SolutionOther1984")
}
