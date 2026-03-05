package p328

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

func TestSolution328(t *testing.T) {
	assert := assert.New(t)
	l := util.InitLinkedList([]int{1, 2, 3, 4, 5})
	res := oddEvenList(l)
	actual := make([]int, 0)
	for res != nil {
		actual = append(actual, res.Val)
		res = res.Next
	}
	assert.Equal([]int{1, 3, 5, 2, 4}, actual, "Solution322")
}
