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

func TestSolution994(t *testing.T) {
	assert := assert.New(t)
	actual := orangesRotting([][]int{{2, 1, 1}, {1, 1, 0}, {0, 1, 1}})
	assert.Equal(actual, 4, "Solution994")
}
