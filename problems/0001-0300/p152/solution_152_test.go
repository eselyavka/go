package p152

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

func TestSolution152(t *testing.T) {
	assert := assert.New(t)
	actual := maxProduct([]int{-2, 3, -4})
	assert.Equal(24, actual, "Solution152")
}
