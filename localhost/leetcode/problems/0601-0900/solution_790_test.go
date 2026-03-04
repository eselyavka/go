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

func TestSolution790(t *testing.T) {
	assert := assert.New(t)
	actual := numTilings(4)
	assert.Equal(actual, 11, "Solution790")
}
