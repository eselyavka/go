package p322

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

func TestSolution322(t *testing.T) {
	assert := assert.New(t)
	actual := coinChange([]int{1, 2, 5}, 11)
	assert.Equal(3, actual, "Solution322")
}
