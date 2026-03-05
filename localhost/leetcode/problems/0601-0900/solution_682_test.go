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

func TestSolution682(t *testing.T) {
	assert := assert.New(t)
	actual := calPoints([]string{"5", "2", "C", "D", "+"})
	assert.Equal(actual, 30, "Solution682")
}
