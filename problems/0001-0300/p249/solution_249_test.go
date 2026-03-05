package p249

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

func TestSolution249(t *testing.T) {
	assert := assert.New(t)
	res := groupStrings([]string{"abc", "bcd", "acef", "xyz", "az", "ba", "a", "z"})
	expected := [][]string{{"a", "z"}, {"abc", "bcd", "xyz"}, {"az", "ba"}, {"acef"}}
	assert.Equal(len(res), len(expected), "Solution249")

	array_equals := make([]bool, 0)
	for _, arr_res := range res {
		exists := false
		for _, arr_exp := range expected {
			if util.StringSlicesEqual(arr_res, arr_exp) {
				exists = true
			}
		}
		array_equals = append(array_equals, exists)
	}
	assert.Equal(array_equals, []bool{true, true, true, true})
}
