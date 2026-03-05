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

func TestSolution1404(t *testing.T) {
	assert := assert.New(t)
	actual := numSteps("1101")
	assert.Equal(actual, 6, "SolutionOther1404")
}
