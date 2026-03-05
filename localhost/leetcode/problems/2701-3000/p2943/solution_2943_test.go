package p2943

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

func TestSolution2943(t *testing.T) {
	assert := assert.New(t)
	actual := maximizeSquareHoleArea(3, 13, []int{2, 4, 3}, []int{4, 6, 7, 12, 10, 13, 2})
	assert.Equal(actual, 9, "SolutionOther2943")
}
