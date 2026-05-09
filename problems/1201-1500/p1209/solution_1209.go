package p1209

type charCount struct {
	char  uint8
	count int
}

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

	stack := make([]charCount, 0)

	i := 0

	for i < len(s) {
		if len(stack) > 0 && s[i] == stack[len(stack)-1].char {
			node := stack[len(stack)-1]
			stack = stack[:len(stack)-1]
			if node.count < k-1 {
				stack = append(stack, charCount{char: node.char, count: node.count + 1})
			}
		} else {
			stack = append(stack, charCount{char: s[i], count: 1})
		}

		i++
	}

	ans := ""
	for _, node := range stack {
		for i := 0; i < node.count; i++ {
			ans += string(node.char)
		}
	}

	return ans
}
