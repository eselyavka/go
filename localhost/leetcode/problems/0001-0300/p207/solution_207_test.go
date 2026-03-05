package p207

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

func TestSolution207(t *testing.T) {
	assert := assert.New(t)
	actual := canFinish(2, [][]int{{1, 0}})
	assert.True(actual, "Solution207")
}
