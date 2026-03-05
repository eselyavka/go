package p187

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

func TestSolution187(t *testing.T) {
	assert := assert.New(t)
	actual := findRepeatedDnaSequences("AAAAACCCCCAAAAACCCCCCAAAAAGGGTTT")
	sort.Strings(actual)
	expected := []string{"AAAAACCCCC", "CCCCCAAAAA"}
	sort.Strings(expected)
	assert.Equal(expected, actual, "Solution187")
}
