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

func TestSolution1099(t *testing.T) {
	assert := assert.New(t)
	actual := twoSumLessThanK([]int{34, 23, 1, 24, 75, 33, 54, 8}, 60)
	assert.Equal(58, actual, "Solution1099")
}
