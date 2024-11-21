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
	actual := twoSum_1(arr, 9)
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
	res := removeDuplicatesArray(actual)
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

	res := insert_708(&head, 2)
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
	l1 := initLinkedList([]int{2, 4, 3})
	l2 := initLinkedList([]int{5, 6, 4})

	ans := addTwoNumbers(l1, l2)

	assert.Equal(ans.Val, 7, "Solution2")
	assert.Equal(ans.Next.Val, 0, "Solution2")
	assert.Equal(ans.Next.Next.Val, 8, "Solution2")
	assert.Nil(ans.Next.Next.Next)
}

func TestSolution5(t *testing.T) {
	assert := assert.New(t)
	res := longestPalindrome("babad")
	assert.Equal(res, "bab", "Solution5")

	res = longestPalindromeDP("babad")
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
	ans := make([]bool, 0)

	for _, arr := range res {
		if intSliceEqual(arr, []int{-1, 0, 1}) {
			ans = append(ans, true)
			continue
		}

		if intSliceEqual(arr, []int{-1, -1, 2}) {
			ans = append(ans, true)
			continue
		}
	}

	assert.Equal([]bool{true, true}, ans, "Solution15")
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

func TestSolution20(t *testing.T) {
	assert := assert.New(t)
	res := isValid("()")
	assert.Equal(res, true, "Solution20")
}

func TestSolution76(t *testing.T) {
	assert := assert.New(t)
	res := minWindow("ADOBECODEBANC", "ABC")
	assert.Equal(res, "BANC", "Solution76")
}

func TestSolution33(t *testing.T) {
	assert := assert.New(t)
	res := search_33([]int{4, 5, 6, 7, 0, 1, 2}, 0)
	assert.Equal(res, 4, "Solution33")
	res = search_33([]int{1, 2, 3, 4, 5, 6, 7}, 3)
	assert.Equal(res, 2, "Solution33")
}

func TestSolution153(t *testing.T) {
	assert := assert.New(t)
	res := findMin([]int{4, 5, 6, 7, 0, 1, 2})
	assert.Equal(res, 0, "Solution153")
	res = findMin([]int{1, 2, 3, 4, 5, 6, 7})
	assert.Equal(res, 1, "Solution153")
}

func TestSolution206(t *testing.T) {
	assert := assert.New(t)
	l1 := initLinkedList([]int{1, 2, 3, 4, 5})

	res := reverseList(l1)
	actual := make([]int, 0)
	for res != nil {
		actual = append(actual, res.Val)
		res = res.Next
	}

	assert.Equal([]int{5, 4, 3, 2, 1}, actual, "Solution206")
}

func TestSolution21(t *testing.T) {
	assert := assert.New(t)
	l1 := initLinkedList([]int{1, 2, 4})
	l2 := initLinkedList([]int{1, 3, 4})

	res := mergeTwoLists(l1, l2)
	actual := make([]int, 0)
	for res != nil {
		actual = append(actual, res.Val)
		res = res.Next
	}

	assert.Equal([]int{1, 1, 2, 3, 4, 4}, actual, "Solution21")
}

func TestSolution19(t *testing.T) {
	assert := assert.New(t)
	l1 := initLinkedList([]int{1, 2, 3, 4, 5})

	res := removeNthFromEnd(l1, 2)
	actual := make([]int, 0)
	for res != nil {
		actual = append(actual, res.Val)
		res = res.Next
	}

	assert.Equal([]int{1, 2, 3, 5}, actual, "Solution19")

	l2 := initLinkedList([]int{1, 2, 3, 4, 5})
	res = removeNthFromEndFast(l2, 2)
	actual = make([]int, 0)
	for res != nil {
		actual = append(actual, res.Val)
		res = res.Next
	}

	assert.Equal([]int{1, 2, 3, 5}, actual, "Solution19")
}

func TestSolution143(t *testing.T) {
	assert := assert.New(t)
	l1 := initLinkedList([]int{1, 2, 3, 4, 5})

	reorderList(l1)
	actual := make([]int, 0)
	for l1 != nil {
		actual = append(actual, l1.Val)
		l1 = l1.Next
	}

	assert.Equal([]int{1, 5, 2, 4, 3}, actual, "Solution143")

	l2 := initLinkedList([]int{1, 2, 3, 4, 5})
	reorderListMem(l2)
	actual = make([]int, 0)
	for l2 != nil {
		actual = append(actual, l2.Val)
		l2 = l2.Next
	}

	assert.Equal([]int{1, 5, 2, 4, 3}, actual, "Solution143")
}

func TestSolution141(t *testing.T) {
	assert := assert.New(t)
	l1 := initLinkedList([]int{1, 2, 3})

	l1.Next.Next.Next = l1

	actual := hasCycle(l1)

	assert.True(actual, "Solution143")

}

func TestSolution98(t *testing.T) {
	assert := assert.New(t)

	root := TreeNode{Val: 5, Left: nil, Right: nil}
	root.Left = &TreeNode{Val: 4, Left: nil, Right: nil}
	root.Right = &TreeNode{Val: 6, Left: nil, Right: nil}
	root.Right.Left = &TreeNode{Val: 3, Left: nil, Right: nil}
	root.Right.Right = &TreeNode{Val: 7, Left: nil, Right: nil}

	actual := isValidBST(&root)

	assert.False(actual, "Solution98")

}

func TestSolution1631(t *testing.T) {
	assert := assert.New(t)

	actual := minimumEffortPath([][]int{{1, 2, 2}, {3, 8, 2}, {5, 3, 5}})

	assert.Equal(2, actual, "Solution1631")
}

func TestSolution226(t *testing.T) {
	assert := assert.New(t)

	root := TreeNode{Val: 4, Left: nil, Right: nil}
	root.Left = &TreeNode{Val: 2, Left: nil, Right: nil}
	root.Left.Left = &TreeNode{Val: 1, Left: nil, Right: nil}
	root.Left.Right = &TreeNode{Val: 3, Left: nil, Right: nil}
	root.Right = &TreeNode{Val: 7, Left: nil, Right: nil}
	root.Right.Left = &TreeNode{Val: 6, Left: nil, Right: nil}
	root.Right.Right = &TreeNode{Val: 9, Left: nil, Right: nil}

	curr := invertTree(&root)

	q := make([]*TreeNode, 0)
	q = append(q, curr)
	var node *TreeNode
	actual := make([]int, 0)

	for len(q) > 0 {
		node = q[0]
		actual = append(actual, node.Val)
		q = q[1:]
		if node.Left != nil {
			q = append(q, node.Left)
		}

		if node.Right != nil {
			q = append(q, node.Right)
		}
	}

	assert.Equal([]int{4, 7, 2, 9, 6, 3, 1}, actual, "Solution226")
}

func TestSolution23(t *testing.T) {
	assert := assert.New(t)
	l1 := initLinkedList([]int{1, 4, 5})
	l2 := initLinkedList([]int{1, 3, 4})
	l3 := initLinkedList([]int{2, 6})

	lists := []*ListNode{l1, l2, l3}

	ans := mergeKLists(lists)

	actual := make([]int, 0)
	for ans != nil {
		actual = append(actual, ans.Val)
		ans = ans.Next
	}

	assert.Equal([]int{1, 1, 2, 3, 4, 4, 5, 6}, actual, "Solution23")

}

func TestSolution230(t *testing.T) {
	assert := assert.New(t)

	root := TreeNode{Val: 3, Left: nil, Right: nil}
	root.Left = &TreeNode{Val: 1, Left: nil, Right: nil}
	root.Left.Right = &TreeNode{Val: 2, Left: nil, Right: nil}
	root.Right = &TreeNode{Val: 4, Left: nil, Right: nil}

	actual := kthSmallest(&root, 1)

	assert.Equal(1, actual, "Solution230")

}

func TestSolution104(t *testing.T) {
	assert := assert.New(t)

	root := TreeNode{Val: 3, Left: nil, Right: nil}
	root.Left = &TreeNode{Val: 9, Left: nil, Right: nil}
	root.Right = &TreeNode{Val: 20, Left: nil, Right: nil}
	root.Right.Left = &TreeNode{Val: 15, Left: nil, Right: nil}
	root.Right.Right = &TreeNode{Val: 7, Left: nil, Right: nil}

	actual := maxDepth(&root)

	assert.Equal(3, actual, "Solution104")

}

func TestSolution100(t *testing.T) {
	assert := assert.New(t)

	root1 := TreeNode{Val: 1, Left: nil, Right: nil}
	root1.Left = &TreeNode{Val: 2, Left: nil, Right: nil}
	root1.Right = &TreeNode{Val: 3, Left: nil, Right: nil}

	root2 := TreeNode{Val: 1, Left: nil, Right: nil}
	root2.Left = &TreeNode{Val: 2, Left: nil, Right: nil}
	root2.Right = &TreeNode{Val: 3, Left: nil, Right: nil}

	actual := isSameTree(&root1, &root2)

	assert.True(actual, "Solution100")

}

func TestSolution572(t *testing.T) {
	assert := assert.New(t)

	root1 := TreeNode{Val: 3, Left: nil, Right: nil}
	root1.Left = &TreeNode{Val: 4, Left: nil, Right: nil}
	root1.Left.Left = &TreeNode{Val: 1, Left: nil, Right: nil}
	root1.Left.Right = &TreeNode{Val: 2, Left: nil, Right: nil}
	root1.Right = &TreeNode{Val: 5, Left: nil, Right: nil}

	root2 := TreeNode{Val: 4, Left: nil, Right: nil}
	root2.Left = &TreeNode{Val: 1, Left: nil, Right: nil}
	root2.Right = &TreeNode{Val: 2, Left: nil, Right: nil}

	actual := isSubtree(&root1, &root2)

	assert.True(actual, "Solution572")

}

func TestSolution235(t *testing.T) {
	assert := assert.New(t)

	root := TreeNode{Val: 6, Left: nil, Right: nil}
	root.Left = &TreeNode{Val: 2, Left: nil, Right: nil}
	root.Left.Left = &TreeNode{Val: 0, Left: nil, Right: nil}
	root.Left.Right = &TreeNode{Val: 4, Left: nil, Right: nil}
	root.Left.Right.Left = &TreeNode{Val: 3, Left: nil, Right: nil}
	root.Left.Right.Right = &TreeNode{Val: 5, Left: nil, Right: nil}
	root.Right = &TreeNode{Val: 8, Left: nil, Right: nil}
	root.Right.Left = &TreeNode{Val: 7, Left: nil, Right: nil}
	root.Right.Right = &TreeNode{Val: 9, Left: nil, Right: nil}

	actual := lowestCommonAncestorTreeNode(&root, root.Left, root.Right)

	assert.Equal(6, actual.Val, "Solution235")
}

func TestSolution102(t *testing.T) {
	assert := assert.New(t)

	root := TreeNode{Val: 6, Left: nil, Right: nil}
	root.Left = &TreeNode{Val: 2, Left: nil, Right: nil}
	root.Left.Left = &TreeNode{Val: 0, Left: nil, Right: nil}
	root.Left.Right = &TreeNode{Val: 4, Left: nil, Right: nil}
	root.Left.Right.Left = &TreeNode{Val: 3, Left: nil, Right: nil}
	root.Left.Right.Right = &TreeNode{Val: 5, Left: nil, Right: nil}
	root.Right = &TreeNode{Val: 8, Left: nil, Right: nil}
	root.Right.Left = &TreeNode{Val: 7, Left: nil, Right: nil}
	root.Right.Right = &TreeNode{Val: 9, Left: nil, Right: nil}

	actual := levelOrder(&root)

	assert.Equal([][]int{{6}, {2, 8}, {0, 4, 7, 9}, {3, 5}}, actual, "Solution102")
}

func TestSolution213(t *testing.T) {
	assert := assert.New(t)
	res := rob2([]int{1, 2, 3, 1})
	assert.Equal(res, 4, "Solution213")
}

func TestSolution1658(t *testing.T) {
	assert := assert.New(t)
	res := minOperations_1658([]int{3, 2, 20, 1, 1, 3}, 10)
	assert.Equal(res, 5, "Solution1658")
}

func TestSolution647(t *testing.T) {
	assert := assert.New(t)
	res := countSubstrings("abc")
	assert.Equal(res, 3, "Solution647")

	res = countSubstrings("aaa")
	assert.Equal(res, 6, "Solution647")
}

func TestSolution2399(t *testing.T) {
	assert := assert.New(t)
	res := checkDistances("abaccb", []int{1, 3, 0, 5, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0})
	assert.True(res, "Solution2399")
}

func TestSolution2476(t *testing.T) {
	assert := assert.New(t)

	root := TreeNode{Val: 6, Left: nil, Right: nil}
	root.Left = &TreeNode{Val: 2, Left: nil, Right: nil}
	root.Left.Left = &TreeNode{Val: 1, Left: nil, Right: nil}
	root.Left.Right = &TreeNode{Val: 4, Left: nil, Right: nil}
	root.Right = &TreeNode{Val: 13, Left: nil, Right: nil}
	root.Right.Left = &TreeNode{Val: 9, Left: nil, Right: nil}
	root.Right.Right = &TreeNode{Val: 15, Left: nil, Right: nil}
	root.Right.Right.Left = &TreeNode{Val: 14, Left: nil, Right: nil}

	actual := closestNodes(&root, []int{2, 5, 16})

	assert.Equal([][]int{{2, 2}, {4, 6}, {15, -1}}, actual, "Solution2476")
}

func TestSolution2482(t *testing.T) {
	assert := assert.New(t)
	actual := onesMinusZeros([][]int{{0, 1, 1}, {1, 0, 1}, {0, 0, 1}})
	assert.Equal([][]int{{0, 0, 4}, {0, 0, 4}, {-2, -2, 2}}, actual, "Solution2482")
}

func TestSolution91(t *testing.T) {
	assert := assert.New(t)
	actual := numDecodings("112")
	assert.Equal(3, actual, "Solution91")
}

func TestSolution2475(t *testing.T) {
	assert := assert.New(t)
	actual := unequalTriplets([]int{4, 4, 2, 4, 3})
	assert.Equal(3, actual, "Solution2475")
}

func TestSolution2442(t *testing.T) {
	assert := assert.New(t)
	actual := countDistinctIntegers([]int{1, 13, 10, 12, 31})
	assert.Equal(6, actual, "Solution2442")
}

func TestSolution2452(t *testing.T) {
	assert := assert.New(t)
	actual := twoEditWords([]string{"word", "note", "ants", "wood"}, []string{"wood", "joke", "moat"})
	assert.Equal([]string{"word", "note", "wood"}, actual, "Solution2452")
}

func TestSolution2483(t *testing.T) {
	assert := assert.New(t)
	actual := bestClosingTime("YYNY")
	assert.Equal(2, actual, "Solution2483")
}

func TestSolution2408(t *testing.T) {
	assert := assert.New(t)
	actual := Constructor_2408([]string{"one", "two", "three"}, []int{2, 3, 1})
	actual.InsertRow("two", []string{"first", "second", "third"})
	assert.Equal("third", actual.SelectCell("two", 1, 3), "Solution2408")
	actual.InsertRow("two", []string{"fourth", "fifth", "sixth"})
	actual.DeleteRow("two", 1)
	assert.Equal("", actual.SelectCell("two", 1, 2), "Solution2408")
	assert.Equal("fifth", actual.SelectCell("two", 2, 2), "Solution2408")
}

func TestSolution2405(t *testing.T) {
	assert := assert.New(t)
	actual := partitionString("abacaba")
	assert.Equal(4, actual, "Solution2405")
}

func TestSolution322(t *testing.T) {
	assert := assert.New(t)
	actual := coinChange([]int{1, 2, 5}, 11)
	assert.Equal(3, actual, "Solution322")
}

func TestSolution328(t *testing.T) {
	assert := assert.New(t)
	l := initLinkedList([]int{1, 2, 3, 4, 5})
	res := oddEvenList(l)
	actual := make([]int, 0)
	for res != nil {
		actual = append(actual, res.Val)
		res = res.Next
	}
	assert.Equal([]int{1, 3, 5, 2, 4}, actual, "Solution322")
}

func TestSolution152(t *testing.T) {
	assert := assert.New(t)
	actual := maxProduct([]int{-2, 3, -4})
	assert.Equal(24, actual, "Solution152")
}

func TestSolution96(t *testing.T) {
	assert := assert.New(t)
	actual := numTrees(5)
	assert.Equal(42, actual, "Solution96")
}

func TestSolution2487(t *testing.T) {
	assert := assert.New(t)
	l := initLinkedList([]int{5, 2, 13, 3, 8})
	res := removeNodes(l)
	actual := make([]int, 0)
	for res != nil {
		actual = append(actual, res.Val)
		res = res.Next
	}
	assert.Equal([]int{13, 8}, actual, "Solution2487")
}

func TestSolution2436(t *testing.T) {
	assert := assert.New(t)
	actual := minimumSplits([]int{12, 6, 3, 14, 8})
	assert.Equal(2, actual, "Solution2436")
}

func TestSolution2486(t *testing.T) {
	assert := assert.New(t)
	actual := appendCharacters("coaching", "coding")
	assert.Equal(4, actual, "Solution2486")
}

func TestSolution2447(t *testing.T) {
	assert := assert.New(t)
	actual := subarrayGCD([]int{9, 3, 1, 2, 6, 3}, 3)
	assert.Equal(4, actual, "Solution2447")
}

func TestSolution2433(t *testing.T) {
	assert := assert.New(t)
	actual := findArray([]int{5, 2, 0, 3, 1})
	assert.Equal([]int{5, 7, 2, 3, 2}, actual, "Solution2433")
}

func TestSolution2384(t *testing.T) {
	assert := assert.New(t)
	actual := largestPalindromic("444947137")
	assert.Equal("7449447", actual, "Solution2384")
}

func TestSolution2370(t *testing.T) {
	assert := assert.New(t)
	actual := longestIdealString("acfgbd", 2)
	assert.Equal(4, actual, "Solution2370")
}

func TestSolution300(t *testing.T) {
	assert := assert.New(t)
	actual := lengthOfLIS([]int{10, 9, 2, 5, 3, 7, 101, 18})
	assert.Equal(4, actual, "Solution300")
}

func TestSolution2470(t *testing.T) {
	assert := assert.New(t)
	actual := subarrayLCM([]int{3, 6, 2, 7, 1}, 6)
	assert.Equal(4, actual, "Solution2470")
}

func TestSolution2453(t *testing.T) {
	assert := assert.New(t)
	actual := destroyTargets([]int{3, 7, 8, 1, 1, 5}, 2)
	assert.Equal(1, actual, "Solution2453")
}

func TestSolution139(t *testing.T) {
	assert := assert.New(t)
	actual := wordBreak("leetcode", []string{"leet", "code"})
	assert.True(actual, "Solution139")
}

func TestSolution109(t *testing.T) {
	assert := assert.New(t)
	l := initLinkedList([]int{-10, -3, 0, 5, 9})
	root := sortedListToBST(l)
	q := []*TreeNode{root}
	actual := make([]int, 0)
	for len(q) > 0 {
		node := q[0]
		q = q[1:]
		actual = append(actual, node.Val)
		if node.Left != nil {
			q = append(q, node.Left)
		}
		if node.Right != nil {
			q = append(q, node.Right)
		}
	}

	assert.Equal([]int{0, -3, 9, -10, 5}, actual, "Solution109")
}

func TestSolution2500(t *testing.T) {
	assert := assert.New(t)
	actual := deleteGreatestValue([][]int{{1, 2, 4}, {3, 3, 1}})
	assert.Equal(8, actual, "Solution139")
}

func TestSolution2498(t *testing.T) {
	assert := assert.New(t)
	actual := maxJump([]int{0, 2, 5, 6, 7})
	assert.Equal(5, actual, "Solution2498")
}

func TestSolution2491(t *testing.T) {
	assert := assert.New(t)
	actual := dividePlayers([]int{3, 2, 5, 1, 3, 4})
	assert.Equal(int64(22), actual, "Solution2491")
}

func TestSolution695(t *testing.T) {
	assert := assert.New(t)
	actual := maxAreaOfIsland([][]int{{0, 0, 1, 0, 0, 0, 0, 1, 0, 0, 0, 0, 0}, {0, 0, 0, 0, 0, 0, 0, 1, 1, 1, 0, 0, 0}, {0, 1, 1, 0, 1, 0, 0, 0, 0, 0, 0, 0, 0}, {0, 1, 0, 0, 1, 1, 0, 0, 1, 0, 1, 0, 0}, {0, 1, 0, 0, 1, 1, 0, 0, 1, 1, 1, 0, 0}, {0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 1, 0, 0}, {0, 0, 0, 0, 0, 0, 0, 1, 1, 1, 0, 0, 0}, {0, 0, 0, 0, 0, 0, 0, 1, 1, 0, 0, 0, 0}})
	assert.Equal(6, actual, "Solution695")
}

func TestSolution463(t *testing.T) {
	assert := assert.New(t)
	actual := islandPerimeter([][]int{{0, 1, 0, 0}, {1, 1, 1, 0}, {0, 1, 0, 0}, {1, 1, 0, 0}})
	assert.Equal(16, actual, "Solution463")
}

func TestSolution2101(t *testing.T) {
	assert := assert.New(t)
	actual := maximumDetonation([][]int{{2, 1, 3}, {6, 1, 4}})
	assert.Equal(2, actual, "Solution2101")
}

func TestSolution945(t *testing.T) {
	assert := assert.New(t)
	actual := minIncrementForUnique([]int{3, 2, 1, 2, 1, 7})
	assert.Equal(6, actual, "Solution945")
}

func TestSolution55(t *testing.T) {
	assert := assert.New(t)
	actual := canJump([]int{2, 3, 1, 1, 4})
	assert.True(actual, "Solution55")
}

func TestSolution252(t *testing.T) {
	assert := assert.New(t)
	actual := canAttendMeetings([][]int{{0, 30}, {5, 10}, {15, 20}})
	assert.False(actual, "Solution252")
}

func TestSolution56(t *testing.T) {
	assert := assert.New(t)
	actual := mergeIntervals([][]int{{1, 3}, {2, 6}, {8, 10}, {15, 18}})
	assert.Equal([][]int{{1, 6}, {8, 10}, {15, 18}}, actual, "Solution56")
}

func TestSolution159(t *testing.T) {
	assert := assert.New(t)
	actual := lengthOfLongestSubstringTwoDistinct("ccaabbb")
	assert.Equal(5, actual, "Solution159")
}

func TestSolution340(t *testing.T) {
	assert := assert.New(t)
	actual := lengthOfLongestSubstringKDistinct("aa", 1)
	assert.Equal(2, actual, "Solution340")
}

func TestSolution57(t *testing.T) {
	assert := assert.New(t)
	actual := insert_57([][]int{{1, 2}, {3, 5}, {6, 7}, {8, 10}, {12, 16}}, []int{4, 8})
	assert.Equal([][]int{{1, 2}, {3, 10}, {12, 16}}, actual, "Solution57")
}

func TestSolution99(t *testing.T) {
	assert := assert.New(t)
	root := TreeNode{Val: 1, Left: nil, Right: nil}
	root.Left = &TreeNode{Val: 3, Left: nil, Right: nil}
	root.Left.Right = &TreeNode{Val: 2, Left: nil, Right: nil}

	recoverTree(&root)

	res := make([]int, 0)
	q := make([]*TreeNode, 0)
	q = append(q, &root)
	for len(q) > 0 {
		node := q[len(q)-1]
		q = q[:len(q)-1]
		res = append(res, node.Val)

		if node.Left != nil {
			q = append(q, node.Left)
		}

		if node.Right != nil {
			q = append(q, node.Right)
		}
	}
	assert.Equal([]int{3, 1, 2}, res, "Solution99")
}

func TestSolution24(t *testing.T) {
	assert := assert.New(t)
	l := initLinkedList([]int{1, 2, 3, 4})
	root := swapPairs(l)
	actual := make([]int, 0)
	for root != nil {
		actual = append(actual, root.Val)
		root = root.Next
	}

	assert.Equal([]int{2, 1, 4, 3}, actual, "Solution24")
}

func TestSolution435(t *testing.T) {
	assert := assert.New(t)
	actual := eraseOverlapIntervals([][]int{{1, 100}, {11, 22}, {1, 11}, {2, 12}})
	assert.Equal(2, actual, "Solution435")
}

func TestSolution46(t *testing.T) {
	assert := assert.New(t)
	actual := permute([]int{1, 2, 3})
	expected := [][]int{{1, 2, 3}, {1, 3, 2}, {2, 1, 3}, {2, 3, 1}, {3, 1, 2}, {3, 2, 1}}
	var flag bool
	for _, a := range actual {
		flag = false
		for _, b := range expected {
			if intSliceEqual(a, b) {
				flag = true
			}
		}
		assert.True(flag, "Solution46")
	}
}

func TestSolution73(t *testing.T) {
	assert := assert.New(t)
	actual := [][]int{{1, 1, 1}, {1, 0, 1}, {1, 1, 1}}
	expected := [][]int{{1, 0, 1}, {0, 0, 0}, {1, 0, 1}}
	setZeroes(actual)
	var flag bool
	for _, a := range actual {
		flag = false
		for _, b := range expected {
			if intSliceEqual(a, b) {
				flag = true
			}
		}
		assert.True(flag, "Solution73")
	}
}

func TestSolution2391(t *testing.T) {
	assert := assert.New(t)
	actual := garbageCollection([]string{"G", "P", "GP", "GG"}, []int{2, 4, 3})
	assert.Equal(21, actual, "Solution2391")
}

func TestSolution1769(t *testing.T) {
	assert := assert.New(t)
	actual := minOperations_1769("001011")
	assert.Equal([]int{11, 8, 5, 4, 3, 4}, actual, "Solution1658")
}

func TestSolution148(t *testing.T) {
	assert := assert.New(t)
	l := initLinkedList([]int{4, 2, 1, 3})
	root := sortList(l)
	actual := make([]int, 0)
	for root != nil {
		actual = append(actual, root.Val)
		root = root.Next
	}

	assert.Equal([]int{1, 2, 3, 4}, actual, "Solution148")
}

func TestSolution253(t *testing.T) {
	assert := assert.New(t)
	actual := minMeetingRooms([][]int{{0, 30}, {5, 10}, {15, 20}})
	assert.Equal(2, actual, "Solution253")
}

func TestSolution1010(t *testing.T) {
	assert := assert.New(t)
	actual := numPairsDivisibleBy60([]int{30, 20, 150, 100, 40})
	assert.Equal(3, actual, "Solution1010")
}

func TestSolution2038(t *testing.T) {
	assert := assert.New(t)
	actual := winnerOfGame("ABBBBBBBAAA")
	assert.False(actual, "Solution2038")
}

func TestSolution1861(t *testing.T) {
	assert := assert.New(t)
	actual := rotateTheBox([][]byte{{'#', '.', '#'}})
	assert.Equal([][]byte{{'.'}, {'#'}, {'#'}}, actual, "Solution1861")
}

func TestSolution1099(t *testing.T) {
	assert := assert.New(t)
	actual := twoSumLessThanK([]int{34, 23, 1, 24, 75, 33, 54, 8}, 60)
	assert.Equal(58, actual, "Solution1099")
}

func TestSolution204(t *testing.T) {
	assert := assert.New(t)
	actual := countPrimes(10)
	assert.Equal(4, actual, "Solution204")
}

func TestSolution1229(t *testing.T) {
	assert := assert.New(t)
	actual := candyCrush([][]int{{1, 3, 5, 5, 2}, {3, 4, 3, 3, 1}, {3, 2, 4, 5, 2}, {2, 4, 4, 5, 5}, {1, 4, 4, 1, 1}})
	assert.Equal([][]int{{1, 3, 0, 0, 0}, {3, 4, 0, 5, 2}, {3, 2, 0, 3, 1}, {2, 4, 0, 5, 2}, {1, 4, 3, 1, 1}}, actual, "Solution723")
}

func TestSolution258(t *testing.T) {
	assert := assert.New(t)
	actual := addDigits(38)
	assert.Equal(2, actual, "Solution258")
}

func TestSolution54(t *testing.T) {
	assert := assert.New(t)
	actual := spiralOrder([][]int{{1, 2, 3}, {4, 5, 6}, {7, 8, 9}})
	assert.Equal([]int{1, 2, 3, 6, 9, 8, 7, 4, 5}, actual, "Solution54")
}

func TestSolution755(t *testing.T) {
	assert := assert.New(t)
	actual := pourWater([]int{2, 1, 1, 2, 1, 2, 2}, 4, 3)
	assert.Equal([]int{2, 2, 2, 3, 2, 2, 2}, actual, "Solution755")
}

func TestSolution1257(t *testing.T) {
	assert := assert.New(t)
	actual := findSmallestRegion([][]string{{"Earth", "North America", "South America"}, {"North America", "United States", "Canada"}, {"United States", "New York", "Boston"}, {"Canada", "Ontario", "Quebec"}, {"South America", "Brazil"}}, "Quebec", "New York")
	assert.Equal("North America", actual, "Solution1257")

}

func TestSolution370(t *testing.T) {
	assert := assert.New(t)
	actual := getModifiedArray(5, [][]int{{1, 3, 2}, {2, 4, 3}, {0, 2, -2}})
	assert.Equal([]int{-2, 0, 3, 5, 3}, actual, "Solution370")
}

func TestSolution910(t *testing.T) {
	assert := assert.New(t)
	actual := smallestRangeII([]int{1, 3, 6}, 3)
	assert.Equal(3, actual, "Solution910")
}

func TestSolution974(t *testing.T) {
	assert := assert.New(t)
	actual := subarraysDivByK([]int{2, -2, 2, -4}, 6)
	assert.Equal(2, actual, "Solution974")
}

func TestSolution490(t *testing.T) {
	assert := assert.New(t)
	actual := hasPath([][]int{{0, 0, 1, 0, 0}, {0, 0, 0, 0, 0}, {0, 0, 0, 1, 0}, {1, 1, 0, 1, 1}, {0, 0, 0, 0, 0}}, []int{0, 4}, []int{4, 4})
	assert.True(actual, "Solution490")
}

func TestSolution1328(t *testing.T) {
	assert := assert.New(t)
	actual := breakPalindrome("abccba")
	assert.Equal("aaccba", actual, "Solution1328")
}

func TestSolution36(t *testing.T) {
	assert := assert.New(t)
	actual := isValidSudoku([][]byte{
		{'5', '3', '.', '.', '7', '.', '.', '.', '.'},
		{'6', '.', '.', '1', '9', '5', '.', '.', '.'},
		{'.', '9', '8', '.', '.', '.', '.', '6', '.'},
		{'8', '.', '.', '.', '6', '.', '.', '.', '3'},
		{'4', '.', '.', '8', '.', '3', '.', '.', '1'},
		{'7', '.', '.', '.', '2', '.', '.', '.', '6'},
		{'.', '6', '.', '.', '.', '.', '2', '8', '.'},
		{'.', '.', '.', '4', '1', '9', '.', '.', '5'},
		{'.', '.', '.', '.', '8', '.', '.', '7', '9'}})
	assert.True(actual, "Solution36")
}

func TestSolution167(t *testing.T) {
	assert := assert.New(t)

	var arr = []int{2, 7, 11, 15}
	actual := twoSum_167(arr, 9)
	var expected = []int{1, 2}

	assert.Equal(actual, expected, "Solution167")
}

func TestSolution42(t *testing.T) {
	assert := assert.New(t)
	actual := trap([]int{0, 1, 0, 2, 1, 0, 1, 3, 2, 1, 2, 1})
	assert.Equal(6, actual, "Solution42")
}

func TestSolution567(t *testing.T) {
	assert := assert.New(t)
	actual := checkInclusion("ab", "eidbaooo")
	assert.True(actual, "Solution567")
}

func TestSolution1680(t *testing.T) {
	assert := assert.New(t)
	actual := concatenatedBinary(3)
	assert.Equal(27, actual, "Solution1680")
}

func TestSolution452(t *testing.T) {
	assert := assert.New(t)
	actual := findMinArrowShots([][]int{{10, 16}, {2, 8}, {1, 6}, {7, 12}})
	assert.Equal(2, actual, "Solution452")
}

func TestSolution2055(t *testing.T) {
	assert := assert.New(t)
	actual := platesBetweenCandles("**|**|***|", [][]int{{2, 5}, {5, 9}})
	assert.Equal([]int{2, 3}, actual, "Solution2055")
}

func TestSolution239(t *testing.T) {
	assert := assert.New(t)
	actual := maxSlidingWindow([]int{1, 3, -1, -3, 5, 3, 6, 7}, 3)
	assert.Equal([]int{3, 3, 5, 5, 6, 7}, actual, "Solution239")
}

func TestSolution739(t *testing.T) {
	assert := assert.New(t)
	actual := dailyTemperatures([]int{73, 74, 75, 71, 69, 72, 76, 73})
	assert.Equal([]int{1, 1, 4, 2, 1, 1, 0, 0}, actual, "Solution739")
}

func TestSolution283(t *testing.T) {
	assert := assert.New(t)
	actual := []int{0, 1, 0, 3, 12}
	moveZeroes(actual)
	assert.Equal([]int{1, 3, 12, 0, 0}, actual, "Solution283")
}

func TestSolution853(t *testing.T) {
	assert := assert.New(t)
	actual := carFleet(12, []int{10, 8, 0, 5, 3}, []int{2, 4, 1, 1, 3})
	assert.Equal(3, actual, "Solution853")
}

func TestSolution64(t *testing.T) {
	assert := assert.New(t)
	actual := minPathSum([][]int{{1, 3, 1}, {1, 5, 1}, {4, 2, 1}})
	assert.Equal(7, actual, "Solution64")
}

func TestSolution1481(t *testing.T) {
	assert := assert.New(t)
	actual := findLeastNumOfUniqueInts([]int{5, 5, 4}, 1)
	assert.Equal(1, actual, "Solution1481")
}

func TestSolution259(t *testing.T) {
	assert := assert.New(t)
	actual := threeSumSmaller([]int{-2, 0, 1, 3}, 2)
	assert.Equal(2, actual, "Solution259")
}

func TestSolution285(t *testing.T) {
	assert := assert.New(t)
	root := TreeNode{Val: 2, Left: nil, Right: nil}
	root.Left = &TreeNode{Val: 1, Left: nil, Right: nil}
	root.Right = &TreeNode{Val: 3, Left: nil, Right: nil}
	actual := inorderSuccessor(&root, root.Left)
	assert.Equal(2, actual.Val, "Solution285")
}

func TestSolution1419(t *testing.T) {
	assert := assert.New(t)
	actual := minNumberOfFrogs("crcoakroak")
	assert.Equal(2, actual, "Solution1419")
}

func TestSolution150(t *testing.T) {
	assert := assert.New(t)
	actual := evalRPN([]string{"2", "1", "+", "3", "*"})
	assert.Equal(9, actual, "Solution150")
}

func TestSolution823(t *testing.T) {
	assert := assert.New(t)
	actual := numFactoredBinaryTrees([]int{2, 4, 5, 10})
	assert.Equal(7, actual, "Solution823")
}

func TestSolution2517(t *testing.T) {
	assert := assert.New(t)
	actual := maximumTastiness([]int{13, 5, 1, 8, 21, 2}, 3)
	assert.Equal(8, actual, "Solution2517")
}

func TestSolution784(t *testing.T) {
	assert := assert.New(t)
	actual := letterCasePermutation("a1b2")
	assert.Equal([]string{"a1b2", "a1B2", "A1b2", "A1B2"}, actual, "Solution784")
}

func TestSolution78(t *testing.T) {
	assert := assert.New(t)
	actual := subsets([]int{1, 2, 3})
	expected := [][]int{{}, {1}, {2}, {3}, {1, 2}, {1, 2, 3}, {1, 3}, {2, 3}}
	for _, left := range expected {
		is_equal := false
		for _, right := range actual {
			if intSliceEqual(left, right) {
				is_equal = true
			}
		}
		assert.True(is_equal, "Solution78")
	}
}

func TestSolution1448(t *testing.T) {
	assert := assert.New(t)
	root := TreeNode{Val: 3, Left: nil, Right: nil}
	root.Left = &TreeNode{Val: 1, Left: nil, Right: nil}
	root.Left.Left = &TreeNode{Val: 3, Left: nil, Right: nil}
	root.Right = &TreeNode{Val: 4, Left: nil, Right: nil}
	root.Right.Left = &TreeNode{Val: 1, Left: nil, Right: nil}
	root.Right.Right = &TreeNode{Val: 5, Left: nil, Right: nil}

	actual := goodNodes(&root)
	assert.Equal(4, actual, "Solution1448")
}

func TestSolution187(t *testing.T) {
	assert := assert.New(t)
	actual := findRepeatedDnaSequences("AAAAACCCCCAAAAACCCCCCAAAAAGGGTTT")
	sort.Strings(actual)
	expected := []string{"AAAAACCCCC", "CCCCCAAAAA"}
	sort.Strings(expected)
	assert.Equal(expected, actual, "Solution187")
}

func TestSolution155(t *testing.T) {
	assert := assert.New(t)
	actual := Constructor_155()
	actual.Push(-2)
	actual.Push(0)
	actual.Push(-3)
	assert.Equal(-3, actual.GetMin(), "Solution155")
	assert.Equal(-3, actual.Pop(), "Solution155")
	assert.Equal(0, actual.Top(), "Solution155")
	assert.Equal(-2, actual.GetMin(), "Solution155")
}

func TestSolution347(t *testing.T) {
	assert := assert.New(t)
	actual := topKFrequent([]int{1, 1, 1, 2, 2, 3}, 2)
	sort.Ints(actual)
	assert.Equal(actual, []int{1, 2}, "Solution347")
}

func TestSolution967(t *testing.T) {
	assert := assert.New(t)
	actual := numsSameConsecDiff(3, 7)
	sort.Ints(actual)
	assert.Equal(actual, []int{181, 292, 707, 818, 929}, "Solution967")
}

func TestSolution22(t *testing.T) {
	assert := assert.New(t)
	actual := generateParenthesis(2)
	sort.Strings(actual)
	expected := []string{"()()", "(())"}
	sort.Strings(expected)
	assert.Equal(actual, expected, "Solution22")
}

func TestSolution543(t *testing.T) {
	assert := assert.New(t)
	root := TreeNode{Val: 1, Left: nil, Right: nil}
	root.Left = &TreeNode{Val: 2, Left: nil, Right: nil}
	root.Left.Left = &TreeNode{Val: 4, Left: nil, Right: nil}
	root.Left.Right = &TreeNode{Val: 5, Left: nil, Right: nil}
	root.Right = &TreeNode{Val: 3, Left: nil, Right: nil}

	actual := diameterOfBinaryTree(&root)
	assert.Equal(3, actual, "Solution543")
}

func TestSolution110(t *testing.T) {
	assert := assert.New(t)
	root := TreeNode{Val: 1, Left: nil, Right: nil}
	root.Left = &TreeNode{Val: 2, Left: nil, Right: nil}
	root.Left.Left = &TreeNode{Val: 4, Left: nil, Right: nil}
	root.Left.Right = &TreeNode{Val: 5, Left: nil, Right: nil}
	root.Right = &TreeNode{Val: 3, Left: nil, Right: nil}

	actual := isBalanced(&root)
	assert.True(actual, "Solution110")
}

func TestSolution540(t *testing.T) {
	assert := assert.New(t)
	actual := singleNonDuplicate([]int{1, 1, 2, 3, 3, 4, 4, 8, 8})
	assert.Equal(actual, 2, "Solution540")
}

func TestSolution199(t *testing.T) {
	assert := assert.New(t)
	root := TreeNode{Val: 1, Left: nil, Right: nil}
	root.Left = &TreeNode{Val: 2, Left: nil, Right: nil}
	root.Left.Right = &TreeNode{Val: 5, Left: nil, Right: nil}
	root.Right = &TreeNode{Val: 3, Left: nil, Right: nil}
	root.Right.Right = &TreeNode{Val: 4, Left: nil, Right: nil}

	actual := rightSideView(&root)
	assert.Equal([]int{1, 3, 4}, actual, "Solution199")
}

func TestSolution1011(t *testing.T) {
	assert := assert.New(t)
	actual := shipWithinDays([]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}, 5)
	assert.Equal(actual, 15, "Solution1011")
}

func TestSolution410(t *testing.T) {
	assert := assert.New(t)
	actual := splitArray([]int{7, 2, 5, 10, 8}, 2)
	assert.Equal(actual, 18, "Solution410")
}

func TestSolution2031(t *testing.T) {
	assert := assert.New(t)
	actual := subarraysWithMoreZerosThanOnes([]int{0, 1, 1, 0, 1})
	assert.Equal(actual, 9, "Solution2031")
}

func TestSolution17(t *testing.T) {
	assert := assert.New(t)
	actual := letterCombinations("23")
	assert.Equal(actual, []string{"ad", "ae", "af", "bd", "be", "bf", "cd", "ce", "cf"}, "Solution17")
}

func TestSolution2244(t *testing.T) {
	assert := assert.New(t)
	actual := minimumRounds([]int{2, 2, 3, 3, 2, 4, 4, 4, 4, 4})
	assert.Equal(actual, 4, "Solution2244")
}

func TestSolution790(t *testing.T) {
	assert := assert.New(t)
	actual := numTilings(4)
	assert.Equal(actual, 11, "Solution790")
}

func TestSolution443(t *testing.T) {
	assert := assert.New(t)
	actual := compress([]byte{'a', 'a', 'b', 'b', 'c', 'c', 'c'})
	assert.Equal(actual, 6, "Solution443")
}

func TestSolution2007(t *testing.T) {
	assert := assert.New(t)
	actual := findOriginalArray([]int{1, 3, 4, 2, 6, 8})
	assert.Equal(actual, []int{1, 3, 4}, "Solution2007")
}

func TestSolution289(t *testing.T) {
	assert := assert.New(t)
	board := [][]int{{1, 1}, {1, 0}}
	gameOfLife(board)
	assert.Equal(board, [][]int{{1, 1}, {1, 1}}, "Solution289")
}

func TestSolution1413(t *testing.T) {
	assert := assert.New(t)
	actual := minStartValue([]int{-3, 2, -3, 4, 2})
	assert.Equal(actual, 5, "Solution1413")
}

func TestSolution1209(t *testing.T) {
	assert := assert.New(t)
	actual := removeDuplicatesString("pbbcggttciiippooaais", 2)
	assert.Equal(actual, "ps", "Solution1209")
}

func TestSolution746(t *testing.T) {
	assert := assert.New(t)
	actual := minCostClimbingStairs([]int{1, 100, 1, 1, 1, 100, 1, 1, 100, 1})
	assert.Equal(actual, 6, "Solution746")
}

func TestSolution926(t *testing.T) {
	assert := assert.New(t)
	actual := minFlipsMonoIncr("00110")
	assert.Equal(actual, 1, "Solution926")
}

func TestSolution703(t *testing.T) {
	assert := assert.New(t)
	actual := Constructor_703(3, []int{4, 5, 8, 2})
	assert.Equal(actual.Add(3), 4, "Solution703")
	assert.Equal(actual.Add(5), 5, "Solution703")
	assert.Equal(actual.Add(10), 5, "Solution703")
	assert.Equal(actual.Add(9), 8, "Solution703")
	assert.Equal(actual.Add(4), 8, "Solution703")
}

func TestSolution215(t *testing.T) {
	assert := assert.New(t)
	actual := findKthLargest([]int{3, 2, 1, 5, 6, 4}, 2)
	assert.Equal(actual, 5, "Solution215")
}

func TestSolution382(t *testing.T) {
	assert := assert.New(t)

	head := initLinkedList([]int{1, 2, 3})
	actual := Constructor_382(head)

	exists := false
	for _, num := range []int{1, 2, 3} {
		if num == actual.GetRandom() {
			exists = true
		}
	}

	assert.True(exists, "Solution382")
}

func TestSolution250(t *testing.T) {
	assert := assert.New(t)
	root := TreeNode{Val: 5, Left: nil, Right: nil}
	root.Left = &TreeNode{Val: 1, Left: nil, Right: nil}
	root.Left.Left = &TreeNode{Val: 5, Left: nil, Right: nil}
	root.Left.Right = &TreeNode{Val: 5, Left: nil, Right: nil}
	root.Right = &TreeNode{Val: 5, Left: nil, Right: nil}
	root.Right.Right = &TreeNode{Val: 5, Left: nil, Right: nil}

	actual := countUnivalSubtrees(&root)

	assert.Equal(4, actual, "Solution250")
}

func TestSolution994(t *testing.T) {
	assert := assert.New(t)
	actual := orangesRotting([][]int{{2, 1, 1}, {1, 1, 0}, {0, 1, 1}})
	assert.Equal(actual, 4, "Solution994")
}

func TestSolution2348(t *testing.T) {
	assert := assert.New(t)
	actual := zeroFilledSubarray([]int{1, 3, 0, 0, 2, 0, 0, 4})
	assert.Equal(actual, int64(6), "Solution2348")
}

func TestSolution274(t *testing.T) {
	assert := assert.New(t)
	actual := hIndex([]int{3, 0, 6, 1, 5})
	assert.Equal(actual, 3, "Solution274")
}

func TestSolution39(t *testing.T) {
	assert := assert.New(t)
	actual := combinationSum([]int{2, 3, 6, 7}, 7)
	assert.True(int2dSliceIsEqual(actual, [][]int{{2, 2, 3}, {7}}), "Solution39")
}

func TestSolution433(t *testing.T) {
	assert := assert.New(t)
	actual := minMutation("AACCGGTT", "AAACGGTA", []string{"AACCGGTA", "AACCGCTA", "AAACGGTA"})
	assert.Equal(actual, 2, "Solution433")
}

func TestSolution1402(t *testing.T) {
	assert := assert.New(t)
	actual := maxSatisfaction([]int{-1, -8, 0, 5, -9})
	assert.Equal(actual, 14, "Solution1402")
}

func TestSolution90(t *testing.T) {
	assert := assert.New(t)
	actual := subsetsWithDup([]int{1, 2, 2})
	expected := [][]int{{}, {1}, {1, 2}, {1, 2, 2}, {2}, {2, 2}}
	assert.Equal(len(actual), len(expected), "Solution90")
	for _, left := range expected {
		is_equal := false
		for _, right := range actual {
			if intSliceEqual(left, right) {
				is_equal = true
			}
		}
		assert.True(is_equal, "Solution90")
	}
}

func TestSolution881(t *testing.T) {
	assert := assert.New(t)
	actual := numRescueBoats([]int{3, 2, 2, 1}, 3)
	assert.Equal(actual, 3, "Solution881")
}

func TestSolution2226(t *testing.T) {
	assert := assert.New(t)
	actual := maximumCandies([]int{5, 8, 6}, 3)
	assert.Equal(actual, 5, "Solution2226")
}

func TestSolution103(t *testing.T) {
	assert := assert.New(t)
	root := TreeNode{Val: 3, Left: nil, Right: nil}
	root.Left = &TreeNode{Val: 9, Left: nil, Right: nil}
	root.Right = &TreeNode{Val: 20, Left: nil, Right: nil}
	root.Right.Left = &TreeNode{Val: 15, Left: nil, Right: nil}
	root.Right.Right = &TreeNode{Val: 7, Left: nil, Right: nil}

	actual := zigzagLevelOrder(&root)

	assert.Equal([][]int{{3}, {20, 9}, {15, 7}}, actual, "Solution103")
}

func TestSolution1254(t *testing.T) {
	assert := assert.New(t)
	actual := closedIsland([][]int{{1, 1, 1, 1, 1, 1, 1, 0}, {1, 0, 0, 0, 0, 1, 1, 0}, {1, 0, 1, 0, 1, 1, 1, 0}, {1, 0, 0, 0, 0, 1, 0, 1}, {1, 1, 1, 1, 1, 1, 1, 0}})
	assert.Equal(actual, 2, "Solution1254")
}

func TestSolution40(t *testing.T) {
	assert := assert.New(t)
	actual := combinationSum2([]int{10, 1, 2, 7, 6, 1, 5}, 8)
	expected := [][]int{{1, 1, 6}, {1, 2, 5}, {1, 7}, {2, 6}}
	assert.Equal(len(actual), len(expected), "Solution40")
	for _, left := range expected {
		is_equal := false
		for _, right := range actual {
			if intSliceEqual(left, right) {
				is_equal = true
			}
		}
		assert.True(is_equal, "Solution40")
	}
}

func TestSolution1040(t *testing.T) {
	assert := assert.New(t)
	actual := numEnclaves([][]int{{0, 0, 0, 0}, {1, 0, 1, 0}, {0, 1, 1, 0}, {0, 0, 0, 0}})
	assert.Equal(actual, 3, "Solution1040")
}

func TestSolution79(t *testing.T) {
	assert := assert.New(t)
	actual := exist([][]byte{{'A', 'B', 'C', 'E'}, {'S', 'F', 'C', 'S'}, {'A', 'D', 'E', 'E'}}, "ABCCED")
	assert.True(actual)
}

func TestSolution2340(t *testing.T) {
	assert := assert.New(t)
	actual := minimumSwaps([]int{3, 4, 5, 5, 3, 1})
	assert.Equal(actual, 6, "Solution2340")
}

func TestSolution2406(t *testing.T) {
	assert := assert.New(t)
	actual := minGroups([][]int{{5, 10}, {6, 8}, {1, 5}, {2, 3}, {1, 10}})
	assert.Equal(actual, 3, "Solution2406")
}

func TestSolution131(t *testing.T) {
	assert := assert.New(t)
	actual := partition("aab")
	expected := [][]string{{"a", "a", "b"}, {"aa", "b"}}
	assert.Equal(len(actual), len(expected), "Solution131")
	for _, left := range expected {
		is_equal := false
		for _, right := range actual {
			if stringSlicesEqual(left, right) {
				is_equal = true
			}
		}
		assert.True(is_equal, "Solution131")
	}
}

func TestSolution71(t *testing.T) {
	assert := assert.New(t)
	actual := simplifyPath("/home/./foo/")
	assert.Equal(actual, "/home/foo", "Solution71")
}

func TestSolution2125(t *testing.T) {
	assert := assert.New(t)
	actual := numberOfBeams([]string{"011001", "000000", "010100", "001000"})
	assert.Equal(actual, 8, "Solution2125")
}

func TestSolution946(t *testing.T) {
	assert := assert.New(t)
	actual := validateStackSequences([]int{1, 2, 3, 4, 5}, []int{4, 5, 3, 2, 1})
	assert.True(actual, "Solution946")
}

func TestSolution1143(t *testing.T) {
	assert := assert.New(t)
	actual := longestCommonSubsequence("abcde", "ace")
	assert.Equal(actual, 3, "Solution1143")
}

func TestSolution516(t *testing.T) {
	assert := assert.New(t)
	actual := longestPalindromeSubseq("bbbab")
	assert.Equal(actual, 4, "Solution516")
}

func TestSolution704(t *testing.T) {
	assert := assert.New(t)
	res := search_704([]int{4, 5, 6, 7, 0, 1, 2}, 0)
	assert.Equal(res, 4, "Solution704")
	res = search_704([]int{1, 2, 3, 4, 5, 6, 7}, 3)
	assert.Equal(res, 2, "Solution704")
}

func TestSolution74(t *testing.T) {
	assert := assert.New(t)
	actual := searchMatrix([][]int{{1, 3, 5, 7}, {10, 11, 16, 20}, {23, 30, 34, 60}}, 3)
	assert.True(actual, "Solution74")
}

func TestSolution1372(t *testing.T) {
	assert := assert.New(t)
	root := TreeNode{Val: 1, Left: nil, Right: nil}
	root.Left = &TreeNode{Val: 1, Left: nil, Right: nil}
	root.Left.Right = &TreeNode{Val: 1, Left: nil, Right: nil}
	root.Left.Right.Left = &TreeNode{Val: 1, Left: nil, Right: nil}
	root.Left.Right.Left.Right = &TreeNode{Val: 1, Left: nil, Right: nil}
	root.Left.Right.Right = &TreeNode{Val: 1, Left: nil, Right: nil}
	root.Right = &TreeNode{Val: 1, Left: nil, Right: nil}

	actual := longestZigZag(&root)

	assert.Equal(4, actual, "Solution1372")
}

func TestSolution63(t *testing.T) {
	assert := assert.New(t)
	actual := uniquePathsWithObstacles([][]int{{0, 0, 0}, {0, 1, 0}, {0, 0, 0}})
	assert.Equal(actual, 2, "Solution63")
}

func TestSolution117(t *testing.T) {
	assert := assert.New(t)
	root := NodeNext{Val: 1, Left: nil, Right: nil, Next: nil}
	root.Left = &NodeNext{Val: 2, Left: nil, Right: nil, Next: nil}
	root.Left.Left = &NodeNext{Val: 4, Left: nil, Right: nil, Next: nil}
	root.Left.Right = &NodeNext{Val: 5, Left: nil, Right: nil, Next: nil}
	root.Right = &NodeNext{Val: 3, Left: nil, Right: nil, Next: nil}
	root.Right.Right = &NodeNext{Val: 7, Left: nil, Right: nil, Next: nil}

	_ = connect(&root)

	q := make([]*NodeNext, 0)
	q = append(q, &root)

	ans := make([]int, 0)

	for len(q) > 0 {
		node := q[0]
		q = q[1:]
		if node.Left != nil {
			q = append(q, node.Left)
		}
		if node.Right != nil {
			q = append(q, node.Right)
		}
		if node.Next != nil {
			ans = append(ans, node.Next.Val)
		}
	}

	assert.Equal([]int{3, 5, 7}, ans, "Solution177")
}

func TestSolution649(t *testing.T) {
	assert := assert.New(t)
	actual := predictPartyVictory("RDD")
	assert.Equal(actual, "Dire", "Solution649")
}

func TestSolution1456(t *testing.T) {
	assert := assert.New(t)
	actual := maxVowels("abciiidef", 3)
	assert.Equal(actual, 3, "Solution1456")
}

func TestSolution59(t *testing.T) {
	assert := assert.New(t)
	actual := generateMatrix(3)
	assert.Equal(actual, [][]int{{1, 2, 3}, {8, 9, 4}, {7, 6, 5}}, "Solution59")
}

func TestSolution2140(t *testing.T) {
	assert := assert.New(t)
	actual := mostPoints([][]int{{3, 2}, {4, 3}, {4, 4}, {2, 5}})
	assert.Equal(actual, int64(5), "Solution2140")
}

func TestSolution1721(t *testing.T) {
	assert := assert.New(t)
	head := initLinkedList([]int{1, 2, 3, 4, 5})
	it := swapNodes(head, 2)
	actual := make([]int, 0)
	for it != nil {
		actual = append(actual, it.Val)
		it = it.Next
	}
	assert.Equal(actual, []int{1, 4, 3, 2, 5}, "Solution1721")
}

func TestSolution2439(t *testing.T) {
	assert := assert.New(t)
	actual := minimizeArrayValue([]int{3, 7, 1, 6})
	assert.Equal(actual, 5, "Solution2439")
}

func TestSolution2130(t *testing.T) {
	assert := assert.New(t)
	head := initLinkedList([]int{5, 4, 2, 1})
	actual := pairSum(head)
	assert.Equal(actual, 6, "Solution2130")
}

func TestSolution27(t *testing.T) {
	assert := assert.New(t)
	actual := removeElement([]int{3, 2, 2, 3}, 3)
	assert.Equal(actual, 2, "Solution27")
}

func TestSolution1929(t *testing.T) {
	assert := assert.New(t)
	actual := getConcatenation([]int{1, 2, 1})
	assert.Equal(actual, []int{1, 2, 1, 1, 2, 1}, "Solution1929")
}

func TestSolution2265(t *testing.T) {
	assert := assert.New(t)
	root := TreeNode{Val: 4, Left: nil, Right: nil}
	root.Left = &TreeNode{Val: 8, Left: nil, Right: nil}
	root.Left.Left = &TreeNode{Val: 0, Left: nil, Right: nil}
	root.Left.Right = &TreeNode{Val: 1, Left: nil, Right: nil}
	root.Right = &TreeNode{Val: 5, Left: nil, Right: nil}
	root.Right.Right = &TreeNode{Val: 6, Left: nil, Right: nil}

	actual := averageOfSubtree(&root)
	assert.Equal(actual, 5, "Solution2265")
}

func TestSolution2785(t *testing.T) {
	assert := assert.New(t)
	actual := sortVowels("lEetcOde")
	assert.Equal(actual, "lEOtcede", "Solution2785")
}

func TestSolution1930(t *testing.T) {
	assert := assert.New(t)
	actual := countPalindromicSubsequence("bbcbaba")
	assert.Equal(actual, 4, "Solution1930")
}

func TestSolution1436(t *testing.T) {
	assert := assert.New(t)
	actual := destCity([][]string{{"London", "New York"}, {"New York", "Lima"}, {"Lima", "Sao Paulo"}})
	assert.Equal(actual, "Sao Paulo", "Solution1436")
}

func TestSolution682(t *testing.T) {
	assert := assert.New(t)
	actual := calPoints([]string{"5", "2", "C", "D", "+"})
	assert.Equal(actual, 30, "Solution682")
}

func TestSolution2501(t *testing.T) {
	assert := assert.New(t)
	actual := longestSquareStreak([]int{4, 3, 6, 16, 8, 2})
	assert.Equal(actual, 3, "Solution2501")
}

func TestSolution1277(t *testing.T) {
	assert := assert.New(t)
	actual := countSquares([][]int{{0, 1, 1, 1}, {1, 1, 1, 1}, {0, 1, 1, 1}})
	assert.Equal(actual, 15, "Solution1277")
}

func TestSolution3163(t *testing.T) {
	assert := assert.New(t)
	actual := compressedString("abcde")
	assert.Equal(actual, "1a1b1c1d1e", "Solution3163")
	actual = compressedString("aaaaaaaaaaaaaabb")
	assert.Equal(actual, "9a5a2b", "Solution3163")
}

func TestSolution2914(t *testing.T) {
	assert := assert.New(t)
	actual := minChanges("1001")
	assert.Equal(actual, 2, "Solution2914")
}

func TestSolution3011(t *testing.T) {
	assert := assert.New(t)
	actual := canSortArray([]int{8, 4, 2, 30, 15})
	assert.True(actual, "Solution3011")
}

func TestSolution796(t *testing.T) {
	assert := assert.New(t)
	actual := rotateString("abcde", "cdeab")
	assert.True(actual, "Solution796")
}

func TestSolution2490(t *testing.T) {
	assert := assert.New(t)
	actual := isCircularSentence("leetcode eats soul")
	assert.True(actual, "Solution2490")
}

func TestSolution2592(t *testing.T) {
	assert := assert.New(t)
	actual := maximizeGreatness([]int{1, 3, 5, 2, 1, 3, 1})
	assert.Equal(actual, 4, "Solution2592")
}

func TestSolution108(t *testing.T) {
	assert := assert.New(t)
	root := sortedArrayToBST([]int{-10, -3, 0, 5, 9})
	actual := binaryTreeBFS(root)
	assert.Equal(actual, []int{0, -10, 5, -3, 9}, "Solution108")
}

func TestSolution2064(t *testing.T) {
	assert := assert.New(t)
	actual := minimizedMaximum(6, []int{11, 6})
	assert.Equal(actual, 3, "Solution2064")
}

func TestSolution1574(t *testing.T) {
	assert := assert.New(t)
	actual := findLengthOfShortestSubarray([]int{1, 2, 3, 10, 4, 2, 3, 5})
	assert.Equal(actual, 3, "Solution1574")
}

func TestSolution1652(t *testing.T) {
	assert := assert.New(t)
	actual := decrypt([]int{5, 7, 1, 4}, 3)
	assert.Equal(actual, []int{12, 10, 16, 13}, "Solution1652")
}

func TestSolution2461(t *testing.T) {
	assert := assert.New(t)
	actual := maximumSubarraySum([]int{1, 5, 4, 2, 9, 9, 9}, 3)
	assert.Equal(actual, int64(15), "Solution2461")
}

func TestSolution2516(t *testing.T) {
	assert := assert.New(t)
	actual := takeCharacters("aabaaaacaabc", 2)
	assert.Equal(actual, 8, "Solution2516")
}

func TestSolution2257(t *testing.T) {
	assert := assert.New(t)
	actual := countUnguarded(4, 6, [][]int{{0, 0}, {1, 1}, {2, 3}}, [][]int{{0, 1}, {2, 2}, {1, 4}})
	assert.Equal(actual, 7, "Solution2257")
}
