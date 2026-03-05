package p763

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

func TestSolution763(t *testing.T) {
	assert := assert.New(t)
	actual := partitionLabels("ababcbacadefegdehijhklij")
	assert.Equal(actual, []int{9, 7, 8}, "Solution763")
}
