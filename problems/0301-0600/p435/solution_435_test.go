package p435

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

func TestSolution435(t *testing.T) {
	assert := assert.New(t)
	actual := eraseOverlapIntervals([][]int{{1, 100}, {11, 22}, {1, 11}, {2, 12}})
	assert.Equal(2, actual, "Solution435")
}
