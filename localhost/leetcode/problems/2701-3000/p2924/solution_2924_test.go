package p2924

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

func TestSolution2924(t *testing.T) {
	assert := assert.New(t)
	actual := findChampion(3, [][]int{{0, 1}, {1, 2}})
	assert.Equal(actual, 0, "Solution2924")
}
