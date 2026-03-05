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

func TestSolution3152(t *testing.T) {
	assert := assert.New(t)
	actual := isArraySpecial([]int{4, 3, 1, 6}, [][]int{{0, 2}, {2, 3}})
	assert.Equal(actual, []bool{false, true}, "Solution3152")
}
