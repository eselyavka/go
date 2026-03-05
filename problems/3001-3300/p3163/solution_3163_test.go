package p3163

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

func TestSolution3163(t *testing.T) {
	assert := assert.New(t)
	actual := compressedString("abcde")
	assert.Equal(actual, "1a1b1c1d1e", "Solution3163")
	actual = compressedString("aaaaaaaaaaaaaabb")
	assert.Equal(actual, "9a5a2b", "Solution3163")
}
