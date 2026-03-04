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

func TestSolution1257(t *testing.T) {
	assert := assert.New(t)
	actual := findSmallestRegion([][]string{{"Earth", "North America", "South America"}, {"North America", "United States", "Canada"}, {"United States", "New York", "Boston"}, {"Canada", "Ontario", "Quebec"}, {"South America", "Brazil"}}, "Quebec", "New York")
	assert.Equal("North America", actual, "Solution1257")

}
