package p1422

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

func TestSolution1422(t *testing.T) {
	assert := assert.New(t)
	actual := maxScore("011101")
	assert.Equal(actual, 5, "Solution1422")
	actual = maxScore("00")
	assert.Equal(actual, 1, "Solution1422")
}
