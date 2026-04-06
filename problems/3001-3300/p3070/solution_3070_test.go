package p3070

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSolution3070(t *testing.T) {
	assert.Equal(t, 4, countSubmatrices([][]int{{7, 6, 3}, {6, 6, 1}}, 18), "example 1")
	assert.Equal(t, 6, countSubmatrices([][]int{{7, 2, 9}, {1, 5, 0}, {2, 6, 6}}, 20), "example 2")
	assert.Equal(t, 1, countSubmatrices([][]int{{5}}, 5), "single cell within limit")
	assert.Equal(t, 0, countSubmatrices([][]int{{5}}, 4), "single cell above limit")
}
