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

func TestSolution22(t *testing.T) {
	assert := assert.New(t)
	actual := generateParenthesis(2)
	sort.Strings(actual)
	expected := []string{"()()", "(())"}
	sort.Strings(expected)
	assert.Equal(actual, expected, "Solution22")
}
