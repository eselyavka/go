package p76

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

func TestSolution76(t *testing.T) {
	assert := assert.New(t)
	res := minWindow("ADOBECODEBANC", "ABC")
	assert.Equal(res, "BANC", "Solution76")
}
