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

func TestSolution1625(t *testing.T) {
	assert := assert.New(t)
	actual := findLexSmallestString("5525", 9, 2)
	assert.Equal(actual, "2050", "Solution1625")
}
