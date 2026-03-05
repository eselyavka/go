package p1390

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

func TestSolution1390(t *testing.T) {
	assert := assert.New(t)
	actual := sumFourDivisors([]int{21, 4, 7})
	assert.Equal(actual, 32, "SolutionOther1390")
}
