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

func TestSolution24(t *testing.T) {
	assert := assert.New(t)
	l := initLinkedList([]int{1, 2, 3, 4})
	root := swapPairs(l)
	actual := make([]int, 0)
	for root != nil {
		actual = append(actual, root.Val)
		root = root.Next
	}

	assert.Equal([]int{2, 1, 4, 3}, actual, "Solution24")
}
