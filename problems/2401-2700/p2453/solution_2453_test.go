package p2453

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

func TestSolution2453(t *testing.T) {
	assert := assert.New(t)
	actual := destroyTargets([]int{3, 7, 8, 1, 1, 5}, 2)
	assert.Equal(1, actual, "Solution2453")
}
