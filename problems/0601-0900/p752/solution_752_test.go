package p752

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

func TestSolution752(t *testing.T) {
	assert := assert.New(t)

	assert.Equal(6, openLock([]string{"0201", "0101", "0102", "1212", "2002"}, "0202"), "Solution752 example")
	assert.Equal(1, openLock([]string{"8888"}, "0009"), "Solution752 single move")
	assert.Equal(-1, openLock([]string{"8887", "8889", "8878", "8898", "8788", "8988", "7888", "9888"}, "8888"), "Solution752 blocked target")
	assert.Equal(-1, openLock([]string{"0000"}, "8888"), "Solution752 blocked start")
	assert.Equal(0, openLock(nil, "0000"), "Solution752 already unlocked")
}
