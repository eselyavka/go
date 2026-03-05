package p39

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

func TestSolution39(t *testing.T) {
	assert := assert.New(t)
	actual := combinationSum([]int{2, 3, 6, 7}, 7)
	assert.True(util.Int2DSliceIsEqual(actual, [][]int{{2, 2, 3}, {7}}), "Solution39")
}
