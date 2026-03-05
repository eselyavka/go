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

func TestSolution382(t *testing.T) {
	assert := assert.New(t)

	head := initLinkedList([]int{1, 2, 3})
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
