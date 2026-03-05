package p2436

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

func TestSolution2436(t *testing.T) {
	assert := assert.New(t)
	actual := minimumSplits([]int{12, 6, 3, 14, 8})
	assert.Equal(2, actual, "Solution2436")
}
