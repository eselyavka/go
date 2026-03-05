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

func TestSolution1480(t *testing.T) {
	assert := assert.New(t)
	var arr = []int{1, 2, 3, 4}
	actual := runningSum(arr)
	assert.Equal(actual, []int{1, 3, 6, 10}, "Solution1480")
}
