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

	exists := false
	for _, num := range []int{1, 2, 3} {
		if num == actual.GetRandom() {
			exists = true
		}
	}

	assert.True(exists, "Solution382")
}
