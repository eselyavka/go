package p393

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

func TestSolution393(t *testing.T) {
	assert := assert.New(t)
	actual := validUtf8([]int{197, 130, 1})
	assert.True(actual, "SolutionOther393")
}
