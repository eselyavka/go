package p1904

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

func TestSolution1904(t *testing.T) {
	assert := assert.New(t)
	actual := numberOfRounds("09:31", "10:14")
	assert.Equal(actual, 1, "Solution1904")
}
