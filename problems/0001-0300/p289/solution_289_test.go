package p289

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

func TestSolution289(t *testing.T) {
	assert := assert.New(t)
	board := [][]int{{1, 1}, {1, 0}}
	gameOfLife(board)
	assert.Equal(board, [][]int{{1, 1}, {1, 1}}, "Solution289")
}
