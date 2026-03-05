package p33

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

func TestSolution33(t *testing.T) {
	assert := assert.New(t)
	res := search([]int{4, 5, 6, 7, 0, 1, 2}, 0)
	assert.Equal(res, 4, "Solution33")
	res = search([]int{1, 2, 3, 4, 5, 6, 7}, 3)
	assert.Equal(res, 2, "Solution33")
}
