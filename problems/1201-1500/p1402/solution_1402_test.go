package p1402

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

func TestSolution1402(t *testing.T) {
	assert := assert.New(t)
	actual := maxSatisfaction([]int{-1, -8, 0, 5, -9})
	assert.Equal(actual, 14, "Solution1402")
}
