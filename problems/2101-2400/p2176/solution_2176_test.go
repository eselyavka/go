package p2176

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

func TestSolution2176(t *testing.T) {
	assert := assert.New(t)
	actual := countPairs([]int{3, 1, 2, 2, 2, 1, 3}, 2)
	assert.Equal(actual, 4, "Solution2176")
}
