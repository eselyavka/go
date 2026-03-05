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

func TestSolution781(t *testing.T) {
	assert := assert.New(t)
	actual := numRabbits([]int{1, 1, 2})
	assert.Equal(actual, 5, "Solution781")
}
