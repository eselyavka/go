package p2370

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

func TestSolution2370(t *testing.T) {
	assert := assert.New(t)
	actual := longestIdealString("acfgbd", 2)
	assert.Equal(4, actual, "Solution2370")
}
