package p703

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

func TestSolution703(t *testing.T) {
	assert := assert.New(t)
	actual := Constructor(3, []int{4, 5, 8, 2})
	assert.Equal(actual.Add(3), 4, "Solution703")
	assert.Equal(actual.Add(5), 5, "Solution703")
	assert.Equal(actual.Add(10), 5, "Solution703")
	assert.Equal(actual.Add(9), 8, "Solution703")
	assert.Equal(actual.Add(4), 8, "Solution703")
}
