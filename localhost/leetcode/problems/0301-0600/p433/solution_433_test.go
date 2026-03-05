package p433

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

func TestSolution433(t *testing.T) {
	assert := assert.New(t)
	actual := minMutation("AACCGGTT", "AAACGGTA", []string{"AACCGGTA", "AACCGCTA", "AAACGGTA"})
	assert.Equal(actual, 2, "Solution433")
}
