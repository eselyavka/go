package p1209

import "github.com/eseliavka/go/util"

/*
class Solution(object):

	def removeDuplicates(self, s, k):
	    """
	    :type s: str
	    :type k: int
	    :rtype: str
	    """
	    if len(s) < k:
	        return s

	    stack = []

	    i = 0

	    while i < len(s):
	        if stack and s[i] == stack[-1][0]:
	            node = stack.pop()
	            if node[1] < k-1:
	                stack.append((node[0], node[1]+1))
	        else:
	            stack.append((s[i], 1))
	        i+=1


	    return "".join([t[0]*t[1] for t in stack])
*/

func removeDuplicatesString(s string, k int) string {
	if len(s) < k {
		return s
	}

	stack := make([]util.TupleChar, 0)

	i := 0

	for i < len(s) {
		if len(stack) > 0 && s[i] == stack[len(stack)-1].C {
			node := stack[len(stack)-1]
			stack = stack[:len(stack)-1]
			if node.Num < k-1 {
				stack = append(stack, util.TupleChar{C: node.C, Num: node.Num + 1})
			}
		} else {
			stack = append(stack, util.TupleChar{C: s[i], Num: 1})
		}

		i++
	}

	ans := ""
	for _, node := range stack {
		for i := 0; i < node.Num; i++ {
			ans += string(node.C)
		}
	}

	return ans
}
