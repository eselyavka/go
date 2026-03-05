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

func TestSolution2843(t *testing.T) {
	assert := assert.New(t)
	actual := countSymmetricIntegers(1, 100)
	assert.Equal(actual, 9, "Solution2843")
}
