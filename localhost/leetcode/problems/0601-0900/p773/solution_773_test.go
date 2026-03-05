package p773

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

func TestSolution773(t *testing.T) {
	assert := assert.New(t)
	actual := slidingPuzzle([][]int{{4, 1, 2}, {5, 0, 3}})
	assert.Equal(actual, 5, "Solution773")
}
