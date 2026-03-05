package p516

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

func TestSolution516(t *testing.T) {
	assert := assert.New(t)
	actual := longestPalindromeSubseq("bbbab")
	assert.Equal(actual, 4, "Solution516")
}
