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

func TestSolution2348(t *testing.T) {
	assert := assert.New(t)
	actual := zeroFilledSubarray([]int{1, 3, 0, 0, 2, 0, 0, 4})
	assert.Equal(actual, int64(6), "Solution2348")
}

func TestSolution2348Ref(t *testing.T) {
	assert := assert.New(t)
	actual := zeroFilledSubarrayRef([]int{1, 3, 0, 0, 2, 0, 0, 4})
	assert.Equal(actual, int64(6), "Solution2348")
}
