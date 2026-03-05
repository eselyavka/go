package p150

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

func TestSolution150(t *testing.T) {
	assert := assert.New(t)
	actual := evalRPN([]string{"2", "1", "+", "3", "*"})
	assert.Equal(9, actual, "Solution150")
}
