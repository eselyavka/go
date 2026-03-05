package p143

import (
	"github.com/stretchr/testify/assert"
	"localhost/leetcode/util"
	"sort"
	"strings"
	"testing"
)

var (
	_ = sort.Strings
	_ = strings.Join
)

func TestSolution143(t *testing.T) {
	assert := assert.New(t)
	l1 := util.InitLinkedList([]int{1, 2, 3, 4, 5})

	reorderList(l1)
	actual := make([]int, 0)
	for l1 != nil {
		actual = append(actual, l1.Val)
		l1 = l1.Next
	}

	assert.Equal([]int{1, 5, 2, 4, 3}, actual, "Solution143")

	l2 := util.InitLinkedList([]int{1, 2, 3, 4, 5})
	reorderListMem(l2)
	actual = make([]int, 0)
	for l2 != nil {
		actual = append(actual, l2.Val)
		l2 = l2.Next
	}

	assert.Equal([]int{1, 5, 2, 4, 3}, actual, "Solution143")
}
