package p1481

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

func TestSolution1481(t *testing.T) {
	assert := assert.New(t)
	actual := findLeastNumOfUniqueInts([]int{5, 5, 4}, 1)
	assert.Equal(1, actual, "Solution1481")
}
