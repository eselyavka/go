package p1415

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

func TestSolution1415(t *testing.T) {
	assert := assert.New(t)
	actual := getHappyString(3, 9)
	assert.Equal(actual, "cab", "Solution1415")
}
