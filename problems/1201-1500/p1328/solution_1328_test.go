package p1328

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

func TestSolution1328(t *testing.T) {
	assert := assert.New(t)
	actual := breakPalindrome("abccba")
	assert.Equal("aaccba", actual, "Solution1328")
}
