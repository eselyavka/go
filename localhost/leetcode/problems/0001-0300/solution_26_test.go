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

func TestSolution26(t *testing.T) {
	assert := assert.New(t)
	actual := []int{1, 1, 2}
	res := removeDuplicatesArray(actual)
	assert.Equal(res, 2, "Solution26")
}
