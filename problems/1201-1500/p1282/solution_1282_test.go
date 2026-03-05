package p1282

import (
	"github.com/eseliavka/go/util"
	"github.com/stretchr/testify/assert"
	"sort"
	"strings"
	"testing"
)

var (
	_ = sort.Strings
	_ = strings.Join
)

func TestSolution1282(t *testing.T) {
	assert := assert.New(t)
	actual := groupThePeople([]int{3, 3, 3, 3, 3, 1, 3})
	assert.True(util.Int2DSliceIsEqual(actual, [][]int{{0, 1, 2}, {5}, {3, 4, 6}}), "Solution1282")
}
