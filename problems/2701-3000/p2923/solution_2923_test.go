package p2923

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

func TestSolution2923(t *testing.T) {
	assert := assert.New(t)
	actual := findChampionEasy([][]int{{0, 0, 1}, {1, 0, 1}, {0, 0, 0}})
	assert.Equal(actual, 1, "Solution2923")
}
