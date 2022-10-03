package solutions

func productExceptSelf(nums []int) []int {
	n := len(nums)
	pleft := make([]int, n)
	pright := make([]int, n)

	pleft[0] = nums[0]
	pright[n-1] = nums[n-1]

	i := 1
	j := n - 2

	for j >= 0 {
		pleft[i] = pleft[i-1] * nums[i]
		pright[j] = pright[j+1] * nums[j]
		i += 1
		j -= 1
	}

	res := make([]int, n)

	i = 0

	for i < n {
		if i == 0 {
			res[i] = pright[i+1]
		} else if i == n-1 {
			res[i] = pleft[i-1]
		} else {
			res[i] = pleft[i-1] * pright[i+1]
		}

		i += 1
	}

	return res
}
