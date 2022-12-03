package solutions

func unequalTriplets(nums []int) int {
	n := len(nums)
	ans := 0
	s := make(map[int]struct{})

	for i := 0; i < n; i++ {
		for j := i + 1; j < n; j++ {
			for k := j + 1; k < n; k++ {
				s[nums[i]] = struct{}{}
				s[nums[j]] = struct{}{}
				s[nums[k]] = struct{}{}
				if len(s) == 3 {
					ans++
				}
				s = make(map[int]struct{})
			}
		}
	}
	return ans
}
