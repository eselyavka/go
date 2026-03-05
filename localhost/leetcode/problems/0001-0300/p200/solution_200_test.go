package p200

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

func TestSolution200(t *testing.T) {
	assert := assert.New(t)
	var grid = [][]byte{
		[]byte("11000"),
		[]byte("11000"),
		[]byte("00100"),
		[]byte("00011")}
	actual := numIslands(grid)
	assert.Equal(actual, 3, "Solution200")
}
