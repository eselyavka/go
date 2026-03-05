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

func TestSolution287(t *testing.T) {
	assert := assert.New(t)

	var arr = []int{3, 3, 4, 2, 2}
	actual := findDuplicate(arr)

	assert.Equal(actual, 2, "Solution287")
}
