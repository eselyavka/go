package p2461

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

func TestSolution2461(t *testing.T) {
	assert := assert.New(t)
	actual := maximumSubarraySum([]int{1, 5, 4, 2, 9, 9, 9}, 3)
	assert.Equal(actual, int64(15), "Solution2461")
}
