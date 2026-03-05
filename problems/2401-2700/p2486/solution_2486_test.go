package p2486

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

func TestSolution2486(t *testing.T) {
	assert := assert.New(t)
	actual := appendCharacters("coaching", "coding")
	assert.Equal(4, actual, "Solution2486")
}
