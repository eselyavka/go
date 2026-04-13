package p1848

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

func TestSolution1848(t *testing.T) {
	assert := assert.New(t)
	actual := getMinDistance([]int{1, 2, 3, 4, 5}, 5, 3)
	assert.Equal(actual, 1, "Solution1848")
}
