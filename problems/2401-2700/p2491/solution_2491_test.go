package p2491

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

func TestSolution2491(t *testing.T) {
	assert := assert.New(t)
	actual := dividePlayers([]int{3, 2, 5, 1, 3, 4})
	assert.Equal(int64(22), actual, "Solution2491")
}
