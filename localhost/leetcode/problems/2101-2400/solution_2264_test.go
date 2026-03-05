package solutions

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

func TestSolution2264(t *testing.T) {
	assert := assert.New(t)
	actual := largestGoodInteger("6777133339")
	assert.Equal(actual, "777", "Solution2264")
}
