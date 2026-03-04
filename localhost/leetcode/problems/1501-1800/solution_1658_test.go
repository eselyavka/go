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

func TestSolution1658(t *testing.T) {
	assert := assert.New(t)
	res := minOperations_1658([]int{3, 2, 20, 1, 1, 3}, 10)
	assert.Equal(res, 5, "Solution1658")
}
