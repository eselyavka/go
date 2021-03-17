package solutions

import (
  "testing"
  "github.com/stretchr/testify/assert"
)

func TestSolution1(t *testing.T) {
  assert := assert.New(t)

  var arr = []int {2, 7, 11, 15}
  actual := twoSum(arr, 9)
  var expected = []int{0, 1}

  assert.Equal(actual, expected, "Solution1")
}

func TestSolution287(t *testing.T) {
  assert := assert.New(t)

  var arr = []int {1, 3, 4, 2, 2}
  actual := findDuplicate(arr)

  assert.Equal(actual, 2, "Solution287")
}
