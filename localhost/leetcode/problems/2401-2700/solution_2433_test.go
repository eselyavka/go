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

func TestSolution2433(t *testing.T) {
	assert := assert.New(t)
	actual := findArray([]int{5, 2, 0, 3, 1})
	assert.Equal([]int{5, 7, 2, 3, 2}, actual, "Solution2433")
}
