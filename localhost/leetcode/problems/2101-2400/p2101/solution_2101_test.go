package p2101

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

func TestSolution2101(t *testing.T) {
	assert := assert.New(t)
	actual := maximumDetonation([][]int{{2, 1, 3}, {6, 1, 4}})
	assert.Equal(2, actual, "Solution2101")
}
