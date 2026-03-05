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

func TestSolution3396(t *testing.T) {
	assert := assert.New(t)
	actual := minimumOperations_3396([]int{1, 2, 3, 4, 2, 3, 3, 5, 7})
	assert.Equal(actual, 2, "Solution3396")
}
