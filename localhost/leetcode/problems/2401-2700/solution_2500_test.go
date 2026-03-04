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

func TestSolution2500(t *testing.T) {
	assert := assert.New(t)
	actual := deleteGreatestValue([][]int{{1, 2, 4}, {3, 3, 1}})
	assert.Equal(8, actual, "Solution139")
}
