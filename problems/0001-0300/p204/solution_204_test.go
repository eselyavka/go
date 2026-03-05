package p204

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

func TestSolution204(t *testing.T) {
	assert := assert.New(t)
	actual := countPrimes(10)
	assert.Equal(4, actual, "Solution204")
}
