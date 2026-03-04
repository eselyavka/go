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

func TestSolution128(t *testing.T) {
	assert := assert.New(t)
	res := longestConsecutive([]int{0, 3, 7, 2, 5, 8, 4, 6, 0, 1})
	assert.Equal(res, 9, "Solution128")
}
