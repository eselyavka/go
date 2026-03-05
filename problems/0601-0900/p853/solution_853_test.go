package p853

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

func TestSolution853(t *testing.T) {
	assert := assert.New(t)
	actual := carFleet(12, []int{10, 8, 0, 5, 3}, []int{2, 4, 1, 1, 3})
	assert.Equal(3, actual, "Solution853")
}
