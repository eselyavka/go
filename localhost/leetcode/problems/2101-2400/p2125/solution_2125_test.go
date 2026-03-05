package p2125

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

func TestSolution2125(t *testing.T) {
	assert := assert.New(t)
	actual := numberOfBeams([]string{"011001", "000000", "010100", "001000"})
	assert.Equal(actual, 8, "Solution2125")
}
