package p206

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

func TestSolution206(t *testing.T) {
	assert := assert.New(t)
	l1 := util.InitLinkedList([]int{1, 2, 3, 4, 5})

	res := reverseList(l1)
	actual := make([]int, 0)
	for res != nil {
		actual = append(actual, res.Val)
		res = res.Next
	}

	assert.Equal([]int{5, 4, 3, 2, 1}, actual, "Solution206")
}
