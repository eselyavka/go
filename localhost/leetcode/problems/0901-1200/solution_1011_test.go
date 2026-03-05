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

func TestSolution1011(t *testing.T) {
	assert := assert.New(t)
	actual := shipWithinDays([]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}, 5)
	assert.Equal(actual, 15, "Solution1011")
}
