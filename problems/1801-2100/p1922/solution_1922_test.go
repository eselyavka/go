package p1922

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

func TestSolution1922(t *testing.T) {
	assert := assert.New(t)
	actual := countGoodNumbers(50)
	assert.Equal(actual, 564908303, "Solution1922")
}
