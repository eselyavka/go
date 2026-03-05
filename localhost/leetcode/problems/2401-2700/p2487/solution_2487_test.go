package p2487

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

func TestSolution2487(t *testing.T) {
	assert := assert.New(t)
	l := util.InitLinkedList([]int{5, 2, 13, 3, 8})
	res := removeNodes(l)
	actual := make([]int, 0)
	for res != nil {
		actual = append(actual, res.Val)
		res = res.Next
	}
	assert.Equal([]int{13, 8}, actual, "Solution2487")
}
