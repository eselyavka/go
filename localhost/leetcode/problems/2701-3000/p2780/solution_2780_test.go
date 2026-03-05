package p2780

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

func TestSolution2780(t *testing.T) {
	assert := assert.New(t)
	actual := minimumIndex([]int{1, 2, 2, 2})
	assert.Equal(actual, 2, "Solution2780")
}
