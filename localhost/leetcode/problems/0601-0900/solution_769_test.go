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

func TestSolution769(t *testing.T) {
	assert := assert.New(t)
	actual := maxChunksToSorted([]int{1, 0, 2, 3, 4})
	assert.Equal(actual, 4, "Solution769")
}
