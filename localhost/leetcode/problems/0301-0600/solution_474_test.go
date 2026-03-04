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

func TestSolution474(t *testing.T) {
	assert := assert.New(t)
	actual := findMaxForm([]string{"10", "0001", "111001", "1", "0"}, 5, 3)
	assert.Equal(actual, 4, "SolutionOther474")
}
