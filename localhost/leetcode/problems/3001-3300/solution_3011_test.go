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

func TestSolution3011(t *testing.T) {
	assert := assert.New(t)
	actual := canSortArray([]int{8, 4, 2, 30, 15})
	assert.True(actual, "Solution3011")
}
