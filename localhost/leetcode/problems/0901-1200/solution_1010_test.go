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

func TestSolution1010(t *testing.T) {
	assert := assert.New(t)
	actual := numPairsDivisibleBy60([]int{30, 20, 150, 100, 40})
	assert.Equal(3, actual, "Solution1010")
}
