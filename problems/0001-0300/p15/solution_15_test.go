package p15

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

func TestSolution15(t *testing.T) {
	assert := assert.New(t)
	res := threeSum([]int{-1, 0, 1, 2, -1, -4})
	ans := make([]bool, 0)

	for _, arr := range res {
		if util.IntSliceEqual(arr, []int{-1, 0, 1}) {
			ans = append(ans, true)
			continue
		}

		if util.IntSliceEqual(arr, []int{-1, -1, 2}) {
			ans = append(ans, true)
			continue
		}
	}

	assert.Equal([]bool{true, true}, ans, "Solution15")
}
