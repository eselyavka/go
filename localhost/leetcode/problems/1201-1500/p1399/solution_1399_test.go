package p1399

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

func TestSolution1399(t *testing.T) {
	assert := assert.New(t)
	actual := countLargestGroup(13)
	assert.Equal(actual, 4, "Solution1399")
}
