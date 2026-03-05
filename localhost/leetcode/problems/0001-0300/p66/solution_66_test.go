package p66

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

func TestSolution66(t *testing.T) {
	assert := assert.New(t)
	actual := plusOne([]int{9, 9, 9})
	var expected = []int{1, 0, 0, 0}
	assert.Equal(actual, expected, "Solution66")
}
