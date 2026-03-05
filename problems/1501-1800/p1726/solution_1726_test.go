package p1726

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

func TestSolution1726(t *testing.T) {
	assert := assert.New(t)
	actual := tupleSameProduct([]int{2, 3, 4, 6, 8, 12})
	assert.Equal(actual, 40, "Solution1726")
}
