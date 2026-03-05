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

func TestSolution1209(t *testing.T) {
	assert := assert.New(t)
	actual := removeDuplicatesString("pbbcggttciiippooaais", 2)
	assert.Equal(actual, "ps", "Solution1209")
}
