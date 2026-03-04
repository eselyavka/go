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

func TestSolution1455(t *testing.T) {
	assert := assert.New(t)
	actual := isPrefixOfWord("i love eating burger", "burg")
	assert.Equal(actual, 4, "Solution1455")
}
