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

func TestSolution1039(t *testing.T) {
	assert := assert.New(t)
	actual := minScoreTriangulation([]int{3, 7, 4, 5})
	assert.Equal(actual, 144, "Solution1039")
}
