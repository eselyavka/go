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

func TestSolution2487(t *testing.T) {
	assert := assert.New(t)
	l := initLinkedList([]int{5, 2, 13, 3, 8})
	res := removeNodes(l)
	actual := make([]int, 0)
	for res != nil {
		actual = append(actual, res.Val)
		res = res.Next
	}
	assert.Equal([]int{13, 8}, actual, "Solution2487")
}
