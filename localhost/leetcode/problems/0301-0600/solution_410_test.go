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

func TestSolution410(t *testing.T) {
	assert := assert.New(t)
	actual := splitArray([]int{7, 2, 5, 10, 8}, 2)
	assert.Equal(actual, 18, "Solution410")
}
