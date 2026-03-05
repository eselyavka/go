package p2452

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

func TestSolution2452(t *testing.T) {
	assert := assert.New(t)
	actual := twoEditWords([]string{"word", "note", "ants", "wood"}, []string{"wood", "joke", "moat"})
	assert.Equal([]string{"word", "note", "wood"}, actual, "Solution2452")
}
