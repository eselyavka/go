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

func TestSolution2130(t *testing.T) {
	assert := assert.New(t)
	head := initLinkedList([]int{5, 4, 2, 1})
	actual := pairSum(head)
	assert.Equal(actual, 6, "Solution2130")
}
