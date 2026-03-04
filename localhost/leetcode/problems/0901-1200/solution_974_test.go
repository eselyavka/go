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

func TestSolution974(t *testing.T) {
	assert := assert.New(t)
	actual := subarraysDivByK([]int{2, -2, 2, -4}, 6)
	assert.Equal(2, actual, "Solution974")
}
