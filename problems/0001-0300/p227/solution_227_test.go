package p227

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

func TestSolution227(t *testing.T) {
	assert := assert.New(t)
	res := calculate("14-3/2")
	assert.Equal(res, 13, "Solution227")
}
