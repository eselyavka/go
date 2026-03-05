package p2537

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

func TestSolution2537(t *testing.T) {
	assert := assert.New(t)
	actual := countGood([]int{3, 1, 4, 3, 2, 2, 4}, 2)
	assert.Equal(actual, int64(4), "Solution2537")
}
