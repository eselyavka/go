package p967

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

func TestSolution967(t *testing.T) {
	assert := assert.New(t)
	actual := numsSameConsecDiff(3, 7)
	sort.Ints(actual)
	assert.Equal(actual, []int{181, 292, 707, 818, 929}, "Solution967")
}
