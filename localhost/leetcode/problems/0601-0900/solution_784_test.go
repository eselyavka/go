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

func TestSolution784(t *testing.T) {
	assert := assert.New(t)
	actual := letterCasePermutation("a1b2")
	assert.Equal([]string{"a1b2", "a1B2", "A1b2", "A1B2"}, actual, "Solution784")
}
