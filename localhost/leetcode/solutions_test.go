package solutions

import (
	"sort"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSolution1(t *testing.T) {
	assert := assert.New(t)

	var arr = []int{2, 7, 11, 15}
	actual := twoSum(arr, 9)
	var expected = []int{0, 1}

	assert.Equal(actual, expected, "Solution1")
}

func TestSolution287(t *testing.T) {
	assert := assert.New(t)

	var arr = []int{3, 3, 4, 2, 2}
	actual := findDuplicate(arr)

	assert.Equal(actual, 2, "Solution287")
}

func TestSolution49(t *testing.T) {
	assert := assert.New(t)
	tests := []struct {
		expected []string
	}{
		{[]string{"bat"}},
		{[]string{"nat", "tan"}},
		{[]string{"ate", "eat", "tea"}},
	}

	var arr = []string{"eat", "tea", "tan", "ate", "nat", "bat"}
	actual := groupAnagrams(arr)
	for _, expected := range tests {
		is_equal := false
		for _, actual := range actual {
			sort.Strings(expected.expected)
			sort.Strings(actual)
			if strings.Join(expected.expected, "") == strings.Join(actual, "") {
				is_equal = true
			}
		}
		assert.True(is_equal, "Solution49")
	}
}

func TestSolution1480(t *testing.T) {
	assert := assert.New(t)
	var arr = []int{1, 2, 3, 4}
	actual := runningSum(arr)
	assert.Equal(actual, []int{1, 3, 6, 10}, "Solution1480")
}

func TestSolution88(t *testing.T) {
	assert := assert.New(t)
	var nums1 = []int{1, 2, 3, 0, 0, 0}
	var m = 3
	var nums2 = []int{2, 5, 6}
	var n = 3
	merge(nums1, m, nums2, n)
	assert.Equal(nums1, []int{1, 2, 2, 3, 5, 6}, "Solution88")
}

func TestSolution200(t *testing.T) {
	assert := assert.New(t)
	var grid = [][]byte{
		[]byte("11000"),
		[]byte("11000"),
		[]byte("00100"),
		[]byte("00011")}
	actual := numIslands(grid)
	assert.Equal(actual, 3, "Solution200")
}

func TestSolution121(t *testing.T) {
	assert := assert.New(t)
	var prices = []int{7, 1, 5, 3, 6, 4}
	actual := maxProfit(prices)
	assert.Equal(actual, 5, "Solution121")
}

func TestSolution53(t *testing.T) {
	assert := assert.New(t)
	var nums = []int{-2, 1, -3, 4, -1, 2, 1, -5, 4}
	actual := maxSubArray(nums)
	assert.Equal(actual, 6, "Solution53")
}

func TestSolution70(t *testing.T) {
	assert := assert.New(t)
	actual := climbStairs(4)
	assert.Equal(actual, 5, "Solution70")
}

func TestSolution118(t *testing.T) {
	assert := assert.New(t)
	actual := generate(5)
	var expected = [][]int{{1}, {1, 1}, {1, 2, 1}, {1, 3, 3, 1}, {1, 4, 6, 4, 1}}
	assert.Equal(actual, expected, "Solution118")
}

func TestSolution66(t *testing.T) {
	assert := assert.New(t)
	actual := plusOne([]int{9, 9, 9})
	var expected = []int{1, 0, 0, 0}
	assert.Equal(actual, expected, "Solution66")
}

func TestSolution26(t *testing.T) {
	assert := assert.New(t)
	actual := []int{1, 1, 2}
	res := removeDuplicates(actual)
	assert.Equal(res, 2, "Solution26")
}

func TestSolution1457(t *testing.T) {
	assert := assert.New(t)
	root := TreeNode{Val: 2, Left: nil, Right: nil}
	root.Left = &TreeNode{Val: 3, Left: nil, Right: nil}
	root.Left.Left = &TreeNode{3, nil, nil}
	root.Left.Right = &TreeNode{1, nil, nil}
	root.Right = &TreeNode{1, nil, nil}
	root.Right.Right = &TreeNode{1, nil, nil}
	assert.Equal(pseudoPalindromicPaths(&root), 2, "Solution1457")
}

func TestSolution169(t *testing.T) {
	assert := assert.New(t)
	actual := []int{3, 2, 3}
	res := majorityElement(actual)
	assert.Equal(res, 3, "Solution169")
}

func TestSolution844(t *testing.T) {
	assert := assert.New(t)
	res := backspaceCompare("ab#c", "ad#c")
	assert.True(res, "Solution844")
}

func TestSolution48(t *testing.T) {
	assert := assert.New(t)
	actual := [][]int{{1, 2, 3}, {4, 5, 6}, {7, 8, 9}}
	expected := [][]int{{7, 4, 1}, {8, 5, 2}, {9, 6, 3}}
	rotate(actual)
	assert.Equal(actual, expected, "Solution48")
}

func TestSolution938(t *testing.T) {
	assert := assert.New(t)
	root := TreeNode{Val: 10, Left: nil, Right: nil}
	root.Left = &TreeNode{Val: 5, Left: nil, Right: nil}
	root.Left.Left = &TreeNode{3, nil, nil}
	root.Left.Right = &TreeNode{7, nil, nil}
	root.Right = &TreeNode{15, nil, nil}
	root.Right.Right = &TreeNode{18, nil, nil}
	assert.Equal(rangeSumBST(&root, 7, 15), 32, "Solution938")
}

func TestSolution680(t *testing.T) {
	assert := assert.New(t)
	res := validPalindrome("abca")
	assert.True(res, "Solution680")
}

func TestSolution1249(t *testing.T) {
	assert := assert.New(t)
	res := minRemoveToMakeValid("lee(t(c)o)de)")
	assert.Equal(res, "lee(t(c)o)de", "Solution1249")
}

func TestSolution1762(t *testing.T) {
	assert := assert.New(t)
	expected := []int{0, 2, 3}
	res := findBuildings([]int{4, 2, 3, 1})
	assert.Equal(res, expected, "Solution1762")
}

func TestSolution408(t *testing.T) {
	assert := assert.New(t)
	res := validWordAbbreviation("internationalization", "i12ip4n")
	assert.False(res, "Solution408")
}

func TestSolution1570(t *testing.T) {
	assert := assert.New(t)
	v1 := Constructor([]int{1, 0, 0, 2, 3})
	v2 := Constructor([]int{0, 3, 0, 4, 0})
	res := v1.dotProduct(v2)
	assert.Equal(res, 8, "Solution1570")
}
