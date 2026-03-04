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

func TestSolution983(t *testing.T) {
	assert := assert.New(t)
	actual := mincostTickets([]int{1, 4, 6, 7, 8, 20}, []int{2, 7, 15})
	assert.Equal(actual, 11, "Solution983")
}
