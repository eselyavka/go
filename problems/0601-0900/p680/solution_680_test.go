package p680

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

func TestSolution680(t *testing.T) {
	assert := assert.New(t)
	res := validPalindrome("abca")
	assert.True(res, "Solution680")
}
