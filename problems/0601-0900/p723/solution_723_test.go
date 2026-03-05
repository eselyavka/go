package p723

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

func TestSolution723(t *testing.T) {
	assert := assert.New(t)
	actual := candyCrush([][]int{{1, 3, 5, 5, 2}, {3, 4, 3, 3, 1}, {3, 2, 4, 5, 2}, {2, 4, 4, 5, 5}, {1, 4, 4, 1, 1}})
	assert.Equal([][]int{{1, 3, 0, 0, 0}, {3, 4, 0, 5, 2}, {3, 2, 0, 3, 1}, {2, 4, 0, 5, 2}, {1, 4, 3, 1, 1}}, actual, "Solution723")
}
