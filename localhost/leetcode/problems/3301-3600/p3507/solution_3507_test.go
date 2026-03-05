package p3507

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

func TestSolution3507(t *testing.T) {
	assert := assert.New(t)
	actual := minimumPairRemoval([]int{5, 2, 3, 1})
	assert.Equal(actual, 2, "SolutionOther3507")
}
