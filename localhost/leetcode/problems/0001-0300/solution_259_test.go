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

func TestSolution259(t *testing.T) {
	assert := assert.New(t)
	actual := threeSumSmaller([]int{-2, 0, 1, 3}, 2)
	assert.Equal(2, actual, "Solution259")
}
