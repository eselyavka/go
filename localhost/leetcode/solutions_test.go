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
	res := search([]int{4, 5, 6, 7, 0, 1, 2}, 0)
	assert.Equal(res, 4, "Solution33")
	res = search([]int{1, 2, 3, 4, 5, 6, 7}, 3)
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
	res := minOperations([]int{3, 2, 20, 1, 1, 3}, 10)
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
	actual := Constructor([]string{"one", "two", "three"}, []int{2, 3, 1})
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

func TestSolution2348(t *testing.T) {
	assert := assert.New(t)
	actual := largestPalindromic("444947137")
	assert.Equal("7449447", actual, "Solution2348")
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
