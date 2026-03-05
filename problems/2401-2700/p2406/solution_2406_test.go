package p2406

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

func TestSolution2406(t *testing.T) {
	assert := assert.New(t)
	actual := minGroups([][]int{{5, 10}, {6, 8}, {1, 5}, {2, 3}, {1, 10}})
	assert.Equal(actual, 3, "Solution2406")
}
