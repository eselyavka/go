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

func TestSolution2439(t *testing.T) {
	assert := assert.New(t)
	actual := minimizeArrayValue([]int{3, 7, 1, 6})
	assert.Equal(actual, 5, "Solution2439")
}
