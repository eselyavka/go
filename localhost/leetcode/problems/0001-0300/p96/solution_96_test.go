package p96

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

func TestSolution96(t *testing.T) {
	assert := assert.New(t)
	actual := numTrees(5)
	assert.Equal(42, actual, "Solution96")
}
