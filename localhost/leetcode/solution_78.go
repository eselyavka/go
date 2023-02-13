package solutions

func rec_78(nums []int, idx int, subset []int, acc *[][]int) {
	temp := make([]int, len(subset), len(subset))
	copy(temp, subset)
	*acc = append(*acc, temp)
	for i := idx; i < len(nums); i++ {
		subset = append(subset, nums[i])
		rec_78(nums, i+1, subset, acc)
		subset = subset[:len(subset)-1]
	}
}

func subsets(nums []int) [][]int {
	var ans [][]int

	rec_78(nums, 0, []int{}, &ans)

	return ans
}
