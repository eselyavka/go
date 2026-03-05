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

func TestSolution213(t *testing.T) {
	assert := assert.New(t)
	res := rob2([]int{1, 2, 3, 1})
	assert.Equal(res, 4, "Solution213")
}
