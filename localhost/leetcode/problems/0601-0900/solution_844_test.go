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

func TestSolution844(t *testing.T) {
	assert := assert.New(t)
	res := backspaceCompare("ab#c", "ad#c")
	assert.True(res, "Solution844")
}
