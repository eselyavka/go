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

func TestSolution1(t *testing.T) {
	assert := assert.New(t)

	var arr = []int{2, 7, 11, 15}
	actual := twoSum_1(arr, 9)
	var expected = []int{0, 1}

	assert.Equal(actual, expected, "Solution1")
}
