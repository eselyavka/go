package p71

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

func TestSolution71(t *testing.T) {
	assert := assert.New(t)
	actual := simplifyPath("/home/./foo/")
	assert.Equal(actual, "/home/foo", "Solution71")
}
