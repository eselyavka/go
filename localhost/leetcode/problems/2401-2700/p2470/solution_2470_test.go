package p2470

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

func TestSolution2470(t *testing.T) {
	assert := assert.New(t)
	actual := subarrayLCM([]int{3, 6, 2, 7, 1}, 6)
	assert.Equal(4, actual, "Solution2470")
}
