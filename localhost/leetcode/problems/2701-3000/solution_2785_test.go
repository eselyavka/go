package solutions

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

func TestSolution2785(t *testing.T) {
	assert := assert.New(t)
	actual := sortVowels("lEetcOde")
	assert.Equal(actual, "lEOtcede", "Solution2785")
}
