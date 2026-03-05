package p3100

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

func TestSolution3100(t *testing.T) {
	assert := assert.New(t)
	actual := maxBottlesDrunk(13, 6)
	assert.Equal(actual, 15, "Solution3100")
}
