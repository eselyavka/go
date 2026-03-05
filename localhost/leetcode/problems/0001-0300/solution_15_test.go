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

func TestSolution15(t *testing.T) {
	assert := assert.New(t)
	res := threeSum([]int{-1, 0, 1, 2, -1, -4})
	ans := make([]bool, 0)

	for _, arr := range res {
		if intSliceEqual(arr, []int{-1, 0, 1}) {
			ans = append(ans, true)
			continue
		}

		if intSliceEqual(arr, []int{-1, -1, 2}) {
			ans = append(ans, true)
			continue
		}
	}

	assert.Equal([]bool{true, true}, ans, "Solution15")
}
