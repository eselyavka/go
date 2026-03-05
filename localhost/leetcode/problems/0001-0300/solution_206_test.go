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

func TestSolution206(t *testing.T) {
	assert := assert.New(t)
	l1 := initLinkedList([]int{1, 2, 3, 4, 5})

	res := reverseList(l1)
	actual := make([]int, 0)
	for res != nil {
		actual = append(actual, res.Val)
		res = res.Next
	}

	assert.Equal([]int{5, 4, 3, 2, 1}, actual, "Solution206")
}
