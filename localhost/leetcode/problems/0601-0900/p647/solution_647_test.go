package p647

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

func TestSolution647(t *testing.T) {
	assert := assert.New(t)
	res := countSubstrings("abc")
	assert.Equal(res, 3, "Solution647")

	res = countSubstrings("aaa")
	assert.Equal(res, 6, "Solution647")
}
