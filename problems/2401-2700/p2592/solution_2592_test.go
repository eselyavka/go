package p2592

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

func TestSolution2592(t *testing.T) {
	assert := assert.New(t)
	actual := maximizeGreatness([]int{1, 3, 5, 2, 1, 3, 1})
	assert.Equal(actual, 4, "Solution2592")
}
