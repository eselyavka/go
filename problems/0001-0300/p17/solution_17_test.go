package p17

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

func TestSolution17(t *testing.T) {
	assert := assert.New(t)
	actual := letterCombinations("23")
	assert.Equal(actual, []string{"ad", "ae", "af", "bd", "be", "bf", "cd", "ce", "cf"}, "Solution17")
}
