package p153

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

func TestSolution153(t *testing.T) {
	assert := assert.New(t)
	res := findMin([]int{4, 5, 6, 7, 0, 1, 2})
	assert.Equal(res, 0, "Solution153")
	res = findMin([]int{1, 2, 3, 4, 5, 6, 7})
	assert.Equal(res, 1, "Solution153")
}
