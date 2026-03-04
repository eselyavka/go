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

func TestSolution258(t *testing.T) {
	assert := assert.New(t)
	actual := addDigits(38)
	assert.Equal(2, actual, "Solution258")
}
