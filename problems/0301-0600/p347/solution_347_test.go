package p347

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

func TestSolution347(t *testing.T) {
	assert := assert.New(t)
	actual := topKFrequent([]int{1, 1, 1, 2, 2, 3}, 2)
	sort.Ints(actual)
	assert.Equal(actual, []int{1, 2}, "Solution347")
}
