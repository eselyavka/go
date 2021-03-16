package solutions

import (
  "testing"
  "github.com/stretchr/testify/assert"
)

func TestSomething(t *testing.T) {
  assert := assert.New(t)

  var arr = []int {2, 7, 11, 15}
  actual := twoSum(arr, 9)
  var expected = []int{0, 1}

  assert.Equal(actual, expected, "The two arrays should be the same.")
}
