package p1861

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

func TestSolution1861(t *testing.T) {
	assert := assert.New(t)
	actual := rotateTheBox([][]byte{{'#', '.', '#'}})
	assert.Equal([][]byte{{'.'}, {'#'}, {'#'}}, actual, "Solution1861")
}
