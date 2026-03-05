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

func TestSolution2340(t *testing.T) {
	assert := assert.New(t)
	actual := minimumSwaps([]int{3, 4, 5, 5, 3, 1})
	assert.Equal(actual, 6, "Solution2340")
}
