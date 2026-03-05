package p2447

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

func TestSolution2447(t *testing.T) {
	assert := assert.New(t)
	actual := subarrayGCD([]int{9, 3, 1, 2, 6, 3}, 3)
	assert.Equal(4, actual, "Solution2447")
}
