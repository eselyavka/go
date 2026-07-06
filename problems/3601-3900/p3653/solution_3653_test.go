package p3653

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSolution3653(t *testing.T) {
	assert := assert.New(t)

	assert.Equal(4, xorAfterQueries([]int{1, 1, 1}, [][]int{{0, 2, 1, 4}}))
	assert.Equal(31, xorAfterQueries([]int{2, 3, 1, 5, 4}, [][]int{{1, 4, 2, 3}, {0, 2, 1, 2}}))
	assert.Equal(999300005, xorAfterQueries([]int{1_000_000_000, 2}, [][]int{{0, 0, 1, 100_000}}))
}
