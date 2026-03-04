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

func TestSolution1436(t *testing.T) {
	assert := assert.New(t)
	actual := destCity([][]string{{"London", "New York"}, {"New York", "Lima"}, {"Lima", "Sao Paulo"}})
	assert.Equal(actual, "Sao Paulo", "Solution1436")
}
