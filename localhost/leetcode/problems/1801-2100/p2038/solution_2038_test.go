package p2038

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

func TestSolution2038(t *testing.T) {
	assert := assert.New(t)
	actual := winnerOfGame("ABBBBBBBAAA")
	assert.False(actual, "Solution2038")
}
