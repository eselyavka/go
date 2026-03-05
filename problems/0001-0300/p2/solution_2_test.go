package p2

import (
	"github.com/eseliavka/go/util"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestSolution2(t *testing.T) {
	l1 := util.InitLinkedList([]int{2, 4, 3})
	l2 := util.InitLinkedList([]int{5, 6, 4})

	ans := addTwoNumbers(l1, l2)

	assert.Equal(t, 7, ans.Val)
	assert.Equal(t, 0, ans.Next.Val)
	assert.Equal(t, 8, ans.Next.Next.Val)
	assert.Nil(t, ans.Next.Next.Next)
}
