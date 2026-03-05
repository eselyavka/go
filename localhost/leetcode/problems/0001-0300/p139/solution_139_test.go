package p139

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

func TestSolution139(t *testing.T) {
	assert := assert.New(t)
	actual := wordBreak("leetcode", []string{"leet", "code"})
	assert.True(actual, "Solution139")
}
