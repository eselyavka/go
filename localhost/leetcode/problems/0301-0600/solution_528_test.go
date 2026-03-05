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

func TestSolution528(t *testing.T) {
	assert := assert.New(t)
	obj := ConstructorPickIndex([]int{1, 3})
	actual := make([]int, 0)
	expected := []int{0, 1, 1, 1, 1}
	for intSliceEqual(actual, expected) != true {
		actual = make([]int, 0)
		for i := 0; i < 5; i++ {
			actual = append(actual, obj.PickIndex())
		}

	}

	assert.Equal(actual, expected, "Solution528")
}
