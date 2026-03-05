package p2516

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

func TestSolution2516(t *testing.T) {
	assert := assert.New(t)
	actual := takeCharacters("aabaaaacaabc", 2)
	assert.Equal(actual, 8, "Solution2516")
}
