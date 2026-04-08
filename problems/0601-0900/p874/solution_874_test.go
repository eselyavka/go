package p874

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSolution874(t *testing.T) {
	assert := assert.New(t)

	assert.Equal(25, robotSim([]int{4, -1, 3}, [][]int{}), "example without obstacles")
	assert.Equal(65, robotSim([]int{4, -1, 4, -2, 4}, [][]int{{2, 4}}), "example with obstacle")
	assert.Equal(36, robotSim([]int{6, -1, -1, 6}, [][]int{}), "u-turn preserves farthest distance")
}
