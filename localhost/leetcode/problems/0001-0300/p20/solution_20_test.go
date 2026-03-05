package p20

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

func TestSolution20(t *testing.T) {
	assert := assert.New(t)
	res := isValid("()")
	assert.Equal(res, true, "Solution20")
}
