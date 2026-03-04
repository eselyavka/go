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

func TestSolution148(t *testing.T) {
	assert := assert.New(t)
	l := initLinkedList([]int{4, 2, 1, 3})
	root := sortList(l)
	actual := make([]int, 0)
	for root != nil {
		actual = append(actual, root.Val)
		root = root.Next
	}

	assert.Equal([]int{1, 2, 3, 4}, actual, "Solution148")
}
