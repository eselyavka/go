package p649

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

func TestSolution649(t *testing.T) {
	assert := assert.New(t)
	actual := predictPartyVictory("RDD")
	assert.Equal(actual, "Dire", "Solution649")
}
