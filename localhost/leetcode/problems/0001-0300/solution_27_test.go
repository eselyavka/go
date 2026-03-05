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

func TestSolution27(t *testing.T) {
	assert := assert.New(t)
	actual := removeElement([]int{3, 2, 2, 3}, 3)
	assert.Equal(actual, 2, "Solution27")
}
