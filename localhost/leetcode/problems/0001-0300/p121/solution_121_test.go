package p121

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

func TestSolution121(t *testing.T) {
	assert := assert.New(t)
	var prices = []int{7, 1, 5, 3, 6, 4}
	actual := maxProfit(prices)
	assert.Equal(actual, 5, "Solution121")
}
