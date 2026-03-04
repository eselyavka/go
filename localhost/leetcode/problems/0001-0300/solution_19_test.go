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

func TestSolution19(t *testing.T) {
	assert := assert.New(t)
	l1 := initLinkedList([]int{1, 2, 3, 4, 5})

	res := removeNthFromEnd(l1, 2)
	actual := make([]int, 0)
	for res != nil {
		actual = append(actual, res.Val)
		res = res.Next
	}

	assert.Equal([]int{1, 2, 3, 5}, actual, "Solution19")

	l2 := initLinkedList([]int{1, 2, 3, 4, 5})
	res = removeNthFromEndFast(l2, 2)
	actual = make([]int, 0)
	for res != nil {
		actual = append(actual, res.Val)
		res = res.Next
	}

	assert.Equal([]int{1, 2, 3, 5}, actual, "Solution19")
}
