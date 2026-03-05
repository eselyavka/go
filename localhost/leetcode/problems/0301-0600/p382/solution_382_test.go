package p382

import (
	"github.com/stretchr/testify/assert"
	"localhost/leetcode/util"
	"sort"
	"strings"
	"testing"
)

var (
	_ = sort.Strings
	_ = strings.Join
)

func TestSolution382(t *testing.T) {
	assert := assert.New(t)

	head := util.InitLinkedList([]int{1, 2, 3})
	actual := Constructor_382(head)

	allowed := map[int]struct{}{
		1: {},
		2: {},
		3: {},
	}

	for i := 0; i < 50; i++ {
		_, ok := allowed[actual.GetRandom()]
		assert.True(ok, "Solution382")
	}
}
