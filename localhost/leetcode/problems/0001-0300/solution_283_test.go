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

func TestSolution283(t *testing.T) {
	assert := assert.New(t)
	actual := []int{0, 1, 0, 3, 12}
	moveZeroes(actual)
	assert.Equal([]int{1, 3, 12, 0, 0}, actual, "Solution283")
}
