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

func TestSolution2661(t *testing.T) {
	assert := assert.New(t)
	actual := firstCompleteIndex([]int{1, 3, 4, 2}, [][]int{{1, 4}, {2, 3}})
	assert.Equal(actual, 2, "Solution2661")
}
