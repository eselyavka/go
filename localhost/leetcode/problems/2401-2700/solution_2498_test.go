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

func TestSolution2498(t *testing.T) {
	assert := assert.New(t)
	actual := maxJump([]int{0, 2, 5, 6, 7})
	assert.Equal(5, actual, "Solution2498")
}
