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

func TestSolution704(t *testing.T) {
	assert := assert.New(t)
	res := search_704([]int{4, 5, 6, 7, 0, 1, 2}, 0)
	assert.Equal(res, 4, "Solution704")
	res = search_704([]int{1, 2, 3, 4, 5, 6, 7}, 3)
	assert.Equal(res, 2, "Solution704")
}
