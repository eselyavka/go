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

func TestSolution252(t *testing.T) {
	assert := assert.New(t)
	actual := canAttendMeetings([][]int{{0, 30}, {5, 10}, {15, 20}})
	assert.False(actual, "Solution252")
}
