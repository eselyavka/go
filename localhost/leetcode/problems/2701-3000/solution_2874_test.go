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

func TestSolution2874(t *testing.T) {
	assert := assert.New(t)
	actual := maximumTripletValueTwo([]int{12, 6, 1, 2, 7})
	assert.Equal(actual, int64(77), "Solution2874")
}
