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

func TestSolution300(t *testing.T) {
	assert := assert.New(t)
	actual := lengthOfLIS([]int{10, 9, 2, 5, 3, 7, 101, 18})
	assert.Equal(4, actual, "Solution300")
}
