package p91

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

func TestSolution91(t *testing.T) {
	assert := assert.New(t)
	actual := numDecodings("112")
	assert.Equal(3, actual, "Solution91")
}
