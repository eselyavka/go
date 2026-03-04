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

func TestSolution271(t *testing.T) {
	assert := assert.New(t)
	var strs = []string{"Hello", "world"}
	var codec Codec
	encode := codec.Encode(strs)
	decode := codec.Decode(encode)
	assert.Equal(decode, strs, "Solution271")
}
