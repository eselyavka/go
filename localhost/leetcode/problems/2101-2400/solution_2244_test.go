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

func TestSolution2244(t *testing.T) {
	assert := assert.New(t)
	actual := minimumRounds([]int{2, 2, 3, 3, 2, 4, 4, 4, 4, 4})
	assert.Equal(actual, 4, "Solution2244")
}
