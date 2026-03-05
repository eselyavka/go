package p340

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

func TestSolution340(t *testing.T) {
	assert := assert.New(t)
	actual := lengthOfLongestSubstringKDistinct("aa", 1)
	assert.Equal(2, actual, "Solution340")
}
