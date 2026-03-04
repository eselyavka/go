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

func TestSolution739(t *testing.T) {
	assert := assert.New(t)
	actual := dailyTemperatures([]int{73, 74, 75, 71, 69, 72, 76, 73})
	assert.Equal([]int{1, 1, 4, 2, 1, 1, 0, 0}, actual, "Solution739")
}
