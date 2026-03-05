package p21

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

func TestSolution21(t *testing.T) {
	assert := assert.New(t)
	l1 := util.InitLinkedList([]int{1, 2, 4})
	l2 := util.InitLinkedList([]int{1, 3, 4})

	res := mergeTwoLists(l1, l2)
	actual := make([]int, 0)
	for res != nil {
		actual = append(actual, res.Val)
		res = res.Next
	}

	assert.Equal([]int{1, 1, 2, 3, 4, 4}, actual, "Solution21")
}
