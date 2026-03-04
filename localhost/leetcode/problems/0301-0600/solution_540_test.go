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

func TestSolution540(t *testing.T) {
	assert := assert.New(t)
	actual := singleNonDuplicate([]int{1, 1, 2, 3, 3, 4, 4, 8, 8})
	assert.Equal(actual, 2, "Solution540")
}
