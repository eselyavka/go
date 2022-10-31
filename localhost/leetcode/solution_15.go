package solutions

import "sort"

func threeSum(nums []int) [][]int {
	res := make(map[tuple3]struct{})
	dups := make(map[int]struct{})
	seen := make(map[int]int)
	n := len(nums)
	for i := 0; i < n; i++ {
		if _, ok := dups[nums[i]]; !ok {
			dups[nums[i]] = struct{}{}
			for j := i + 1; j < n; j++ {
				s := (nums[i] + nums[j]) * -1
				if val, ok := seen[s]; ok {
					if val == i {
						ints := []int{nums[i], nums[j], s}
						sort.Ints(ints)
						ans := tuple3{num1: ints[0], num2: ints[1], num3: ints[2]}
						res[ans] = struct{}{}
					}
				}
				seen[nums[j]] = i
			}
		}
	}

	ans := make([][]int, len(res))
	i := 0
	for k, _ := range res {
		ans[i] = []int{k.num1, k.num2, k.num3}
		i += 1
	}

	return ans
}
