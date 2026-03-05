package p2483

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

func TestSolution2483(t *testing.T) {
	assert := assert.New(t)
	actual := bestClosingTime("YYNY")
	assert.Equal(2, actual, "Solution2483")
}
