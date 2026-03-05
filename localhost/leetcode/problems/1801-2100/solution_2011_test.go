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

func TestSolution2011(t *testing.T) {
	assert := assert.New(t)
	actual := finalValueAfterOperations([]string{"--X", "X++", "X++"})
	assert.Equal(actual, 1, "Solution2011")
}
