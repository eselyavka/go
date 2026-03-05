package p1518

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

func TestSolution1518(t *testing.T) {
	assert := assert.New(t)
	actual := numWaterBottles(15, 4)
	assert.Equal(actual, 19, "Solution1518")
}
