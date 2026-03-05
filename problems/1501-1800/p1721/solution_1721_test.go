package p1721

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

func TestSolution1721(t *testing.T) {
	assert := assert.New(t)
	head := util.InitLinkedList([]int{1, 2, 3, 4, 5})
	it := swapNodes(head, 2)
	actual := make([]int, 0)
	for it != nil {
		actual = append(actual, it.Val)
		it = it.Next
	}
	assert.Equal(actual, []int{1, 4, 3, 2, 5}, "Solution1721")
}
