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

func TestSolution1980(t *testing.T) {
	assert := assert.New(t)
	actual := findDifferentBinaryString([]string{"01", "10"})
	set := make(map[string]bool)
	for _, str := range []string{"00", "11"} {
		set[str] = true
	}
	assert.True(set[actual], "Solution1980")
}
