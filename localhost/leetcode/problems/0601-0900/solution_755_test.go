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

func TestSolution755(t *testing.T) {
	assert := assert.New(t)
	actual := pourWater([]int{2, 1, 1, 2, 1, 2, 2}, 4, 3)
	assert.Equal([]int{2, 2, 2, 3, 2, 2, 2}, actual, "Solution755")
}
