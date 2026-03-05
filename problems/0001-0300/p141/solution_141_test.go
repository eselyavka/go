package p141

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

func TestSolution141(t *testing.T) {
	assert := assert.New(t)
	l1 := util.InitLinkedList([]int{1, 2, 3})

	l1.Next.Next.Next = l1

	actual := hasCycle(l1)

	assert.True(actual, "Solution143")

}
