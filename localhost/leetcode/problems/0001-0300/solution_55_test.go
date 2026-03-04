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

func TestSolution55(t *testing.T) {
	assert := assert.New(t)
	actual := canJump([]int{2, 3, 1, 1, 4})
	assert.True(actual, "Solution55")
}
