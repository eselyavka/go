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

func TestSolution198(t *testing.T) {
	assert := assert.New(t)
	res := rob([]int{2, 7, 9, 3, 1})
	assert.Equal(res, 12, "Solution198")
}
