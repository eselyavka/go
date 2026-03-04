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

func TestSolution141(t *testing.T) {
	assert := assert.New(t)
	l1 := initLinkedList([]int{1, 2, 3})

	l1.Next.Next.Next = l1

	actual := hasCycle(l1)

	assert.True(actual, "Solution143")

}
