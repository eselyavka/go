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

func TestSolution1570(t *testing.T) {
	assert := assert.New(t)
	v1 := ConstructorSparseVector([]int{1, 0, 0, 2, 3})
	v2 := ConstructorSparseVector([]int{0, 3, 0, 4, 0})
	res := v1.dotProduct(v2)
	assert.Equal(res, 8, "Solution1570")
}
