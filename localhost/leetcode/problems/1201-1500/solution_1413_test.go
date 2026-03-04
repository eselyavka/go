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

func TestSolution1413(t *testing.T) {
	assert := assert.New(t)
	actual := minStartValue([]int{-3, 2, -3, 4, 2})
	assert.Equal(actual, 5, "Solution1413")
}
