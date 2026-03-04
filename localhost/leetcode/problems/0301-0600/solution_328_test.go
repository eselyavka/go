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

func TestSolution328(t *testing.T) {
	assert := assert.New(t)
	l := initLinkedList([]int{1, 2, 3, 4, 5})
	res := oddEvenList(l)
	actual := make([]int, 0)
	for res != nil {
		actual = append(actual, res.Val)
		res = res.Next
	}
	assert.Equal([]int{1, 3, 5, 2, 4}, actual, "Solution322")
}
