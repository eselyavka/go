package p1680

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

func TestSolution1680(t *testing.T) {
	assert := assert.New(t)
	actual := concatenatedBinary(3)
	assert.Equal(27, actual, "Solution1680")
}
