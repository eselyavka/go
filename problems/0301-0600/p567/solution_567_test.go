package p567

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

func TestSolution567(t *testing.T) {
	assert := assert.New(t)
	actual := checkInclusion("ab", "eidbaooo")
	assert.True(actual, "Solution567")
}
