package p1669

import (
	"sort"
	"strings"
	"testing"

	"github.com/eseliavka/go/util"
	"github.com/stretchr/testify/assert"
)

var (
	_ = sort.Strings
	_ = strings.Join
)

func TestSolution1669(t *testing.T) {
	assert := assert.New(t)

	list1 := util.InitLinkedList([]int{0, 1, 2, 3, 4, 5})
	list2 := util.InitLinkedList([]int{1000000, 1000001, 1000002})

	res := mergeInBetween(list1, 3, 4, list2)
	actual := make([]int, 0)
	for res != nil {
		actual = append(actual, res.Val)
		res = res.Next
	}

	assert.Equal([]int{0, 1, 2, 1000000, 1000001, 1000002, 5}, actual, "Solution1669")
}
