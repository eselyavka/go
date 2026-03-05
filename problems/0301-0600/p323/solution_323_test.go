package p323

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

func TestSolution323(t *testing.T) {
	assert := assert.New(t)
	actual := countComponents(5, [][]int{{0, 1}, {1, 2}, {3, 4}})
	assert.Equal(actual, 2, "Solution323")
}
