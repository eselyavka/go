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

func TestSolution11(t *testing.T) {
	assert := assert.New(t)
	res := maxArea([]int{1, 8, 6, 2, 5, 4, 8, 3, 7})
	assert.Equal(res, 49, "Solution11")
}
