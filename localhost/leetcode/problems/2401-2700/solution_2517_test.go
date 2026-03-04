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

func TestSolution2517(t *testing.T) {
	assert := assert.New(t)
	actual := maximumTastiness([]int{13, 5, 1, 8, 21, 2}, 3)
	assert.Equal(8, actual, "Solution2517")
}
