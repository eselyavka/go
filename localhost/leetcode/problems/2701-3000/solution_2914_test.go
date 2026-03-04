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

func TestSolution2914(t *testing.T) {
	assert := assert.New(t)
	actual := minChanges("1001")
	assert.Equal(actual, 2, "Solution2914")
}
