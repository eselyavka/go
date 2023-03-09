package solutions

import "sort"

/*
class Solution(object):

	def findOriginalArray(self, changed):
	    """
	    :type changed: List[int]
	    :rtype: List[int]
	    """
	    n = len(changed)

	    if n % 2 != 0:
	        return []

	    freq = Counter(changed)
	    changed.sort()

	    ans = []
	    for num in changed:
	        if num in freq and freq[num] > 0:
	            freq[num] -= 1
	            double_num = 2 * num
	            if double_num in freq and freq[double_num] > 0:
	                ans.append(num)
	                freq[double_num] -= 1
	            else:
	                return []

	    return ans
*/
func findOriginalArray(changed []int) []int {
	n := len(changed)

	if n%2 != 0 {
		return []int{}
	}

	freq := make(map[int]int)

	for _, num := range changed {
		freq[num]++
	}

	sort.Ints(changed)

	ans := make([]int, 0)

	for _, num := range changed {
		if _, ok := freq[num]; ok && freq[num] > 0 {
			freq[num]--
			double_num := 2 * num
			if _, ok := freq[double_num]; ok && freq[double_num] > 0 {
				ans = append(ans, num)
				freq[double_num]--
			} else {
				return []int{}
			}
		}
	}

	return ans
}
