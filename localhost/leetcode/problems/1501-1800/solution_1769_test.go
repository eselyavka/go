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

func TestSolution1769(t *testing.T) {
	assert := assert.New(t)
	actual := minOperations_1769("001011")
	assert.Equal([]int{11, 8, 5, 4, 3, 4}, actual, "Solution1658")
}
