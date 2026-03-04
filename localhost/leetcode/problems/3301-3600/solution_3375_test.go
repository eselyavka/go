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

func TestSolution3375(t *testing.T) {
	assert := assert.New(t)
	actual := minOperations_3375([]int{5, 2, 5, 4, 5}, 2)
	assert.Equal(actual, 2, "Solution3375")
}
