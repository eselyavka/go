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

func TestSolution2033(t *testing.T) {
	assert := assert.New(t)
	actual := minOperations_2033([][]int{{2, 4}, {6, 8}}, 2)
	assert.Equal(actual, 4, "Solution2033")
}
