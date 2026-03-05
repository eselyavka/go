package p274

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

func TestSolution274(t *testing.T) {
	assert := assert.New(t)
	actual := hIndex([]int{3, 0, 6, 1, 5})
	assert.Equal(actual, 3, "Solution274")
}
