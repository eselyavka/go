package p3392

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

func TestSolution3392(t *testing.T) {
	assert := assert.New(t)
	actual := countSubarrays([]int{1, 2, 1, 4, 1})
	assert.Equal(actual, 1, "Solution3392")
}
