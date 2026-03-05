package p3542

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

func TestSolution3542(t *testing.T) {
	assert := assert.New(t)
	actual := minOperations_3542([]int{1, 2, 1, 2, 1, 2})
	assert.Equal(actual, 4, "SolutionOther3542")
}
