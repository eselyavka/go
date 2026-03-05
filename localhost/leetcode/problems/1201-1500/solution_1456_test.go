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

func TestSolution1456(t *testing.T) {
	assert := assert.New(t)
	actual := maxVowels("abciiidef", 3)
	assert.Equal(actual, 3, "Solution1456")
}
