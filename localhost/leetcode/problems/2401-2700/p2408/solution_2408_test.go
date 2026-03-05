package p2408

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

func TestSolution2408(t *testing.T) {
	assert := assert.New(t)
	actual := Constructor_2408([]string{"one", "two", "three"}, []int{2, 3, 1})
	actual.InsertRow("two", []string{"first", "second", "third"})
	assert.Equal("third", actual.SelectCell("two", 1, 3), "Solution2408")
	actual.InsertRow("two", []string{"fourth", "fifth", "sixth"})
	actual.DeleteRow("two", 1)
	assert.Equal("", actual.SelectCell("two", 1, 2), "Solution2408")
	assert.Equal("fifth", actual.SelectCell("two", 2, 2), "Solution2408")
}
