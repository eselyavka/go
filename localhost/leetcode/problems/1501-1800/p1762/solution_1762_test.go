package p1762

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

func TestSolution1762(t *testing.T) {
	assert := assert.New(t)
	expected := []int{0, 2, 3}
	res := findBuildings([]int{4, 2, 3, 1})
	assert.Equal(res, expected, "Solution1762")
}
