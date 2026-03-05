package p42

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

func TestSolution42(t *testing.T) {
	assert := assert.New(t)
	actual := trap([]int{0, 1, 0, 2, 1, 0, 1, 3, 2, 1, 2, 1})
	assert.Equal(6, actual, "Solution42")
}
