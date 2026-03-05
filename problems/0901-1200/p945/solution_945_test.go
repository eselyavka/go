package p945

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

func TestSolution945(t *testing.T) {
	assert := assert.New(t)
	actual := minIncrementForUnique([]int{3, 2, 1, 2, 1, 7})
	assert.Equal(6, actual, "Solution945")
}
