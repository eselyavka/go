package p2064

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

func TestSolution2064(t *testing.T) {
	assert := assert.New(t)
	actual := minimizedMaximum(6, []int{11, 6})
	assert.Equal(actual, 3, "Solution2064")
}
