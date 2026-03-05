package p408

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

func TestSolution408(t *testing.T) {
	assert := assert.New(t)
	res := validWordAbbreviation("internationalization", "i12ip4n")
	assert.False(res, "Solution408")
}
