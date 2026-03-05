package solutions

/*
class Solution(object):
    def longestOnes(self, nums, k):
        """
        :type nums: List[int]
        :type k: int
        :rtype: int
        """
        n = len(nums)
        right = 0
        left = 0
        ans = 0

        while right < n:
            if nums[right] == 1 or nums[right] == 0 and k > 0:
                if nums[right] == 0:
                    k -= 1
                right += 1
            else:
                ans = max(ans, right - left)
                while k == 0:
                    if nums[left] == 0: k += 1
                    left += 1

            ans = max(ans, right - left)

        return ans
*/

func longestOnes(nums []int, k int) int {
	n := len(nums)
	right := 0
	left := 0
	ans := 0

	for right < n {
		if nums[right] == 1 || (nums[right] == 0 && k > 0) {
			if nums[right] == 0 {
				k--
			}
			right++
		} else {
			ans = MaxInts([]int{ans, right - left})
			for k == 0 {
				if nums[left] == 0 {
					k++
				}
				left++
			}
		}
		ans = MaxInts([]int{ans, right - left})
	}

	return ans
}
