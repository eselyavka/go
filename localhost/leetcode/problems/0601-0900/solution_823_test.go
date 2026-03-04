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

func TestSolution823(t *testing.T) {
	assert := assert.New(t)
	actual := numFactoredBinaryTrees([]int{2, 4, 5, 10})
	assert.Equal(7, actual, "Solution823")
}
