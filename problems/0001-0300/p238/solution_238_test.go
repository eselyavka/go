package p238

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

func TestSolution238(t *testing.T) {
	assert := assert.New(t)
	res := productExceptSelf([]int{1, 2, 3, 4})
	assert.Equal(res, []int{24, 12, 8, 6}, "Solution238")
}
