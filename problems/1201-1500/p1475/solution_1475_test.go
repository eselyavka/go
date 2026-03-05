package p1475

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

func TestSolution1475(t *testing.T) {
	assert := assert.New(t)
	actual := finalPrices([]int{8, 4, 6, 2, 3})
	assert.Equal(actual, []int{4, 2, 4, 2, 3}, "Solution1475")
}
