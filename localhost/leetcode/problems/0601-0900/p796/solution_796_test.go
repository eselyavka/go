package p796

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

func TestSolution796(t *testing.T) {
	assert := assert.New(t)
	actual := rotateString("abcde", "cdeab")
	assert.True(actual, "Solution796")
}
