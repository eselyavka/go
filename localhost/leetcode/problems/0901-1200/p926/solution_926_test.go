package p926

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

func TestSolution926(t *testing.T) {
	assert := assert.New(t)
	actual := minFlipsMonoIncr("00110")
	assert.Equal(actual, 1, "Solution926")
}
