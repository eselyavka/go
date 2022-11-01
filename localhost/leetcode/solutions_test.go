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
	v1 := ConstructorSparseVector([]int{1, 0, 0, 2, 3})
	v2 := ConstructorSparseVector([]int{0, 3, 0, 4, 0})
	res := v1.dotProduct(v2)
	assert.Equal(res, 8, "Solution1570")
}

func TestSolution1650(t *testing.T) {
	assert := assert.New(t)
	root := Node{Val: 3, Left: nil, Right: nil, Parent: nil}
	root.Left = &Node{Val: 5, Left: nil, Right: nil, Parent: &root}
	root.Left.Left = &Node{Val: 6, Left: nil, Right: nil, Parent: root.Left}
	root.Left.Right = &Node{Val: 2, Left: nil, Right: nil, Parent: root.Left}
	root.Left.Right.Left = &Node{Val: 7, Left: nil, Right: nil, Parent: root.Left.Right}
	root.Left.Right.Right = &Node{Val: 4, Left: nil, Right: nil, Parent: root.Left.Right}
	root.Right = &Node{Val: 1, Left: nil, Right: nil, Parent: &root}
	root.Right.Left = &Node{Val: 0, Left: nil, Right: nil, Parent: root.Right}
	root.Right.Right = &Node{Val: 8, Left: nil, Right: nil, Parent: root.Right}

	p := root.Left
	q := root.Left.Right.Right

	assert.Equal(lowestCommonAncestor(p, q), root.Left, "Solution1650")
}

func TestSolution314(t *testing.T) {
	assert := assert.New(t)
	root := TreeNode{Val: 3, Left: nil, Right: nil}
	root.Left = &TreeNode{Val: 9, Left: nil, Right: nil}
	root.Left.Left = &TreeNode{Val: 4, Left: nil, Right: nil}
	root.Left.Right = &TreeNode{Val: 0, Left: nil, Right: nil}
	root.Right = &TreeNode{Val: 8, Left: nil, Right: nil}
	root.Right.Left = &TreeNode{Val: 1, Left: nil, Right: nil}
	root.Right.Right = &TreeNode{Val: 7, Left: nil, Right: nil}

	assert.Equal(verticalOrder(&root), [][]int{{4}, {9}, {3, 0, 1}, {8}, {7}}, "Solution314")
}

func TestSolution876(t *testing.T) {
	assert := assert.New(t)
	head := ListNode{Val: 1, Next: nil}
	head.Next = &ListNode{Val: 2, Next: nil}
	head.Next.Next = &ListNode{Val: 3, Next: nil}
	head.Next.Next.Next = &ListNode{Val: 4, Next: nil}
	head.Next.Next.Next.Next = &ListNode{Val: 5, Next: nil}

	assert.Equal(middleNode(&head), head.Next.Next, "Solution876")
}

func TestSolution198(t *testing.T) {
	assert := assert.New(t)
	res := rob([]int{2, 7, 9, 3, 1})
	assert.Equal(res, 12, "Solution198")
}

func TestSolution94(t *testing.T) {
	assert := assert.New(t)
	root := TreeNode{Val: 1, Left: nil, Right: nil}
	root.Right = &TreeNode{Val: 2, Left: nil, Right: nil}
	root.Right.Left = &TreeNode{Val: 3, Left: nil, Right: nil}

	assert.Equal(inorderTraversal(&root), []int{1, 3, 2}, "Solution94")
}

func TestSolution114(t *testing.T) {
	assert := assert.New(t)
	root := TreeNode{Val: 1, Left: nil, Right: nil}
	root.Left = &TreeNode{Val: 2, Left: nil, Right: nil}
	root.Left.Left = &TreeNode{Val: 3, Left: nil, Right: nil}
	root.Left.Right = &TreeNode{Val: 4, Left: nil, Right: nil}
	root.Right = &TreeNode{Val: 5, Left: nil, Right: nil}
	root.Right.Right = &TreeNode{Val: 6, Left: nil, Right: nil}

	flatten(&root)

	res := []int{root.Val}
	node := root.Right
	for node != nil {
		res = append(res, node.Val)
		node = node.Right
	}

	assert.Equal(res, []int{1, 2, 3, 4, 5, 6}, "Solution114")
}

func TestSolution249(t *testing.T) {
	assert := assert.New(t)
	res := groupStrings([]string{"abc", "bcd", "acef", "xyz", "az", "ba", "a", "z"})
	expected := [][]string{{"a", "z"}, {"abc", "bcd", "xyz"}, {"az", "ba"}, {"acef"}}
	assert.Equal(len(res), len(expected), "Solution249")

	array_equals := make([]bool, 0)
	for _, arr_res := range res {
		exists := false
		for _, arr_exp := range expected {
			if stringSlicesEqual(arr_res, arr_exp) {
				exists = true
			}
		}
		array_equals = append(array_equals, exists)
	}
	assert.Equal(array_equals, []bool{true, true, true, true})
}

func TestSolution528(t *testing.T) {
	assert := assert.New(t)
	obj := ConstructorPickIndex([]int{1, 3})
	actual := make([]int, 0)
	expected := []int{0, 1, 1, 1, 1}
	for intSliceEqual(actual, expected) != true {
		actual = make([]int, 0)
		for i := 0; i < 5; i++ {
			actual = append(actual, obj.PickIndex())
		}

	}

	assert.Equal(actual, expected, "Solution528")
}

func TestSolution426(t *testing.T) {
	assert := assert.New(t)
	root := Node{Val: 4, Left: nil, Right: nil, Parent: nil}
	root.Left = &Node{Val: 2, Left: nil, Right: nil, Parent: nil}
	root.Left.Left = &Node{Val: 1, Left: nil, Right: nil, Parent: nil}
	root.Left.Right = &Node{Val: 3, Left: nil, Right: nil, Parent: nil}
	root.Right = &Node{Val: 5, Left: nil, Right: nil, Parent: nil}

	res := treeToDoublyList(&root)
	var head *Node = nil

	actual := make([]int, 0)

	for head == nil || res.Val != head.Val {
		if head == nil {
			head = res
		}
		actual = append(actual, res.Val)
		res = res.Right
	}

	assert.Equal(actual, []int{1, 2, 3, 4, 5}, "Solution426")
}

func TestSolution1091(t *testing.T) {
	assert := assert.New(t)
	res := shortestPathBinaryMatrix([][]int{{0, 1}, {1, 0}})
	assert.Equal(res, 2, "Solution1091")
}

func TestSolution498(t *testing.T) {
	assert := assert.New(t)
	res := findDiagonalOrder([][]int{{1, 2, 3}, {4, 5, 6}, {7, 8, 9}})
	assert.Equal(res, []int{1, 2, 4, 7, 5, 3, 6, 8, 9}, "Solution498")
}

func TestSolution227(t *testing.T) {
	assert := assert.New(t)
	res := calculate("14-3/2")
	assert.Equal(res, 13, "Solution227")
}

func TestSolution708(t *testing.T) {
	assert := assert.New(t)
	head := ListNode{Val: 3, Next: nil}
	head.Next = &ListNode{Val: 4, Next: nil}
	head.Next.Next = &ListNode{Val: 1, Next: nil}
	head.Next.Next.Next = &head

	res := insert(&head, 2)
	acc := make([]int, 0)

	curr := res.Next
	prev := res

	for curr != res {
		acc = append(acc, prev.Val)
		prev = curr
		curr = curr.Next
	}

	acc = append(acc, prev.Val)

	assert.Equal(acc, []int{3, 4, 1, 2}, "Solution708")
}

func TestSolution3(t *testing.T) {
	assert := assert.New(t)
	res := lengthOfLongestSubstring("abcabcbb")
	assert.Equal(res, 3, "Solution3")
}

func TestSolution2(t *testing.T) {
	assert := assert.New(t)
	l1 := ListNode{Val: 2, Next: nil}
	l1.Next = &ListNode{Val: 4, Next: nil}
	l1.Next.Next = &ListNode{Val: 3, Next: nil}

	l2 := ListNode{Val: 5, Next: nil}
	l2.Next = &ListNode{Val: 6, Next: nil}
	l2.Next.Next = &ListNode{Val: 4, Next: nil}

	ans := addTwoNumbers(&l1, &l2)

	assert.Equal(ans.Val, 7, "Solution2")
	assert.Equal(ans.Next.Val, 0, "Solution2")
	assert.Equal(ans.Next.Next.Val, 8, "Solution2")
	assert.Nil(ans.Next.Next.Next)
}

func TestSolution5(t *testing.T) {
	assert := assert.New(t)
	res := longestPalindrome("babad")
	assert.Equal(res, "bab", "Solution5")
}

func TestSolution4(t *testing.T) {
	assert := assert.New(t)
	res := findMedianSortedArrays([]int{1, 2}, []int{3})
	assert.Equal(res, 2.0, "Solution4")
}

func TestSolution238(t *testing.T) {
	assert := assert.New(t)
	res := productExceptSelf([]int{1, 2, 3, 4})
	assert.Equal(res, []int{24, 12, 8, 6}, "Solution238")
}

func TestSolution271(t *testing.T) {
	assert := assert.New(t)
	var strs = []string{"Hello", "world"}
	var codec Codec
	encode := codec.Encode(strs)
	decode := codec.Decode(encode)
	assert.Equal(decode, strs, "Solution271")
}

func TestSolution128(t *testing.T) {
	assert := assert.New(t)
	res := longestConsecutive([]int{0, 3, 7, 2, 5, 8, 4, 6, 0, 1})
	assert.Equal(res, 9, "Solution128")
}

func TestSolution15(t *testing.T) {
	assert := assert.New(t)
	res := threeSum([]int{-1, 0, 1, 2, -1, -4})
	assert.Equal(res, [][]int{{-1, 0, 1}, {-1, -1, 2}}, "Solution15")
}

func TestSolution11(t *testing.T) {
	assert := assert.New(t)
	res := maxArea([]int{1, 8, 6, 2, 5, 4, 8, 3, 7})
	assert.Equal(res, 49, "Solution11")
}

func TestSolution424(t *testing.T) {
	assert := assert.New(t)
	res := characterReplacement("AABABBA", 1)
	assert.Equal(res, 4, "Solution424")
}
