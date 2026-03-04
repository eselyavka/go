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

func TestSolution2482(t *testing.T) {
	assert := assert.New(t)
	actual := onesMinusZeros([][]int{{0, 1, 1}, {1, 0, 1}, {0, 0, 1}})
	assert.Equal([][]int{{0, 0, 4}, {0, 0, 4}, {-2, -2, 2}}, actual, "Solution2482")
}
