package p2221

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

func TestSolution2221(t *testing.T) {
	assert := assert.New(t)
	actual := triangularSum([]int{1, 2, 3, 4, 5})
	assert.Equal(actual, 8, "Solution2221")
}
