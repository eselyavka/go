package p424

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

func TestSolution424(t *testing.T) {
	assert := assert.New(t)
	res := characterReplacement("AABABBA", 1)
	assert.Equal(res, 4, "Solution424")
}
