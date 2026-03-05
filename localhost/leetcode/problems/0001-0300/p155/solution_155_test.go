package p155

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

func TestSolution155(t *testing.T) {
	assert := assert.New(t)
	actual := Constructor()
	actual.Push(-2)
	actual.Push(0)
	actual.Push(-3)
	assert.Equal(-3, actual.GetMin(), "Solution155")
	assert.Equal(-3, actual.Pop(), "Solution155")
	assert.Equal(0, actual.Top(), "Solution155")
	assert.Equal(-2, actual.GetMin(), "Solution155")
}
