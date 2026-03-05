package p239

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

func TestSolution239(t *testing.T) {
	assert := assert.New(t)
	actual := maxSlidingWindow([]int{1, 3, -1, -3, 5, 3, 6, 7}, 3)
	assert.Equal([]int{3, 3, 5, 5, 6, 7}, actual, "Solution239")
}
