package p2007

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

func TestSolution2007(t *testing.T) {
	assert := assert.New(t)
	actual := findOriginalArray([]int{1, 3, 4, 2, 6, 8})
	assert.Equal(actual, []int{1, 3, 4}, "Solution2007")
}
