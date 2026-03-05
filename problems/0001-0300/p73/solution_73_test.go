package p73

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

func TestSolution73(t *testing.T) {
	assert := assert.New(t)
	actual := [][]int{{1, 1, 1}, {1, 0, 1}, {1, 1, 1}}
	expected := [][]int{{1, 0, 1}, {0, 0, 0}, {1, 0, 1}}
	setZeroes(actual)
	var flag bool
	for _, a := range actual {
		flag = false
		for _, b := range expected {
			if util.IntSliceEqual(a, b) {
				flag = true
			}
		}
		assert.True(flag, "Solution73")
	}
}
