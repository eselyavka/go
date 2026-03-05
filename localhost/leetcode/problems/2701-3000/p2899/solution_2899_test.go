package p2899

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

func TestSolution2899(t *testing.T) {
	assert := assert.New(t)
	actual := lastVisitedIntegers([]int{1, 2, -1, -1, -1})
	assert.Equal(actual, []int{2, 1, -1}, "Solution2899")
}
