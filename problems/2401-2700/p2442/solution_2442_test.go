package p2442

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

func TestSolution2442(t *testing.T) {
	assert := assert.New(t)
	actual := countDistinctIntegers([]int{1, 13, 10, 12, 31})
	assert.Equal(6, actual, "Solution2442")
}
