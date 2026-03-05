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

func TestSolution215(t *testing.T) {
	assert := assert.New(t)
	actual := findKthLargest([]int{3, 2, 1, 5, 6, 4}, 2)
	assert.Equal(actual, 5, "Solution215")
}
