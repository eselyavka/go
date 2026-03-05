package p1652

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

func TestSolution1652(t *testing.T) {
	assert := assert.New(t)
	actual := decrypt([]int{5, 7, 1, 4}, 3)
	assert.Equal(actual, []int{12, 10, 16, 13}, "Solution1652")
}
