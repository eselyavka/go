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

func TestSolution2799(t *testing.T) {
	assert := assert.New(t)
	actual := countCompleteSubarrays([]int{1, 3, 1, 2, 2})
	assert.Equal(actual, 4, "Solution2799")
}
