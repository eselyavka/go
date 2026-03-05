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

func TestSolution2391(t *testing.T) {
	assert := assert.New(t)
	actual := garbageCollection([]string{"G", "P", "GP", "GG"}, []int{2, 4, 3})
	assert.Equal(21, actual, "Solution2391")
}
