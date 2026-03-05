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

func TestSolution2(t *testing.T) {
	assert := assert.New(t)
	l1 := initLinkedList([]int{2, 4, 3})
	l2 := initLinkedList([]int{5, 6, 4})

	ans := addTwoNumbers(l1, l2)

	assert.Equal(ans.Val, 7, "Solution2")
	assert.Equal(ans.Next.Val, 0, "Solution2")
	assert.Equal(ans.Next.Next.Val, 8, "Solution2")
	assert.Nil(ans.Next.Next.Next)
}
