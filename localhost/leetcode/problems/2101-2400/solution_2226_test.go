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

func TestSolution2226(t *testing.T) {
	assert := assert.New(t)
	actual := maximumCandies([]int{5, 8, 6}, 3)
	assert.Equal(actual, 5, "Solution2226")
}
