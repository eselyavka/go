package p443

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

func TestSolution443(t *testing.T) {
	assert := assert.New(t)
	actual := compress([]byte{'a', 'a', 'b', 'b', 'c', 'c', 'c'})
	assert.Equal(actual, 6, "Solution443")
}
