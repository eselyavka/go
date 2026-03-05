package p1419

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

func TestSolution1419(t *testing.T) {
	assert := assert.New(t)
	actual := minNumberOfFrogs("crcoakroak")
	assert.Equal(2, actual, "Solution1419")
}
