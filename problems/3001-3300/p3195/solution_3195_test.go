package p3195

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

func TestSolution3195(t *testing.T) {
	assert := assert.New(t)
	actual := minimumArea([][]int{{0, 1, 0}, {1, 0, 1}})
	assert.Equal(actual, 6, "Solution3195")
}
