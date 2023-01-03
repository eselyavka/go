package solutions

func backtracking_46(nums []int, l, r int, acc *[][]int) {
	if l == r {
		arr := make([]int, 0)
		for _, num := range nums {
			arr = append(arr, num)
		}
		*(acc) = append(*(acc), arr)
	} else {
		for i := l; i < r; i++ {
			buf := nums[l]
			nums[l] = nums[i]
			nums[i] = buf
			backtracking_46(nums, l+1, r, acc)
			buf = nums[l]
			nums[l] = nums[i]
			nums[i] = buf
		}
	}
}

func permute(nums []int) [][]int {
	ans := make([][]int, 0)
	n := len(nums)
	backtracking_46(nums, 0, n, &ans)

	return ans
}
