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

func TestSolution1249(t *testing.T) {
	assert := assert.New(t)
	res := minRemoveToMakeValid("lee(t(c)o)de)")
	assert.Equal(res, "lee(t(c)o)de", "Solution1249")
}
