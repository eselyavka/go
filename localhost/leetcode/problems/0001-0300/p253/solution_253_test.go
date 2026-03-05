package p253

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

func TestSolution253(t *testing.T) {
	assert := assert.New(t)
	actual := minMeetingRooms([][]int{{0, 30}, {5, 10}, {15, 20}})
	assert.Equal(2, actual, "Solution253")
}
