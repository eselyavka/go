package p2130

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

func TestSolution2130(t *testing.T) {
	assert := assert.New(t)
	head := util.InitLinkedList([]int{5, 4, 2, 1})
	actual := pairSum(head)
	assert.Equal(actual, 6, "Solution2130")
}
