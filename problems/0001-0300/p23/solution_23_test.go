package p23

import (
	"github.com/eseliavka/go/util"
	"github.com/stretchr/testify/assert"
	"sort"
	"strings"
	"testing"
)

var (
	_ = sort.Strings
	_ = strings.Join
)

func TestSolution23(t *testing.T) {
	assert := assert.New(t)
	l1 := util.InitLinkedList([]int{1, 4, 5})
	l2 := util.InitLinkedList([]int{1, 3, 4})
	l3 := util.InitLinkedList([]int{2, 6})

	lists := []*util.ListNode{l1, l2, l3}

	ans := mergeKLists(lists)

	actual := make([]int, 0)
	for ans != nil {
		actual = append(actual, ans.Val)
		ans = ans.Next
	}

	assert.Equal([]int{1, 1, 2, 3, 4, 4, 5, 6}, actual, "Solution23")

}
