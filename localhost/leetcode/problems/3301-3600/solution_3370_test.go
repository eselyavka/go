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

func TestSolution3370(t *testing.T) {
	assert := assert.New(t)
	actual := smallestNumber(5)
	assert.Equal(actual, 7, "SolutionOther3370")
}
