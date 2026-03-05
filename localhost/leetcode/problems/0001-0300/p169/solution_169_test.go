package p169

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

func TestSolution169(t *testing.T) {
	assert := assert.New(t)
	actual := []int{3, 2, 3}
	res := majorityElement(actual)
	assert.Equal(res, 3, "Solution169")
}
