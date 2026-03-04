package solutions

func runningSum(nums []int) []int {
	var res []int
	prev := 0
	for _, num := range nums {
		res = append(res, prev+num)
		prev = prev + num
	}
	return res
}
