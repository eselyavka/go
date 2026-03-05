package p881

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

func TestSolution881(t *testing.T) {
	assert := assert.New(t)
	actual := numRescueBoats([]int{3, 2, 2, 1}, 3)
	assert.Equal(actual, 3, "Solution881")
}
