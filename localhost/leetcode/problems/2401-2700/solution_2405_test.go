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

func TestSolution2405(t *testing.T) {
	assert := assert.New(t)
	actual := partitionString("abacaba")
	assert.Equal(4, actual, "Solution2405")
}
