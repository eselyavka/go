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

func TestSolution416(t *testing.T) {
	assert := assert.New(t)
	actual := canPartition([]int{1, 5, 11, 5})
	assert.True(actual, "Solution416")
}
