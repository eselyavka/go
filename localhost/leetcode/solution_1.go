package solutions

func twoSum(nums []int, target int) []int {
	var res []int
	dict := make(map[int]int)
	for i, num := range nums {
		n, ok := dict[target-num]
		if ok {
			res = append(res, n, i)
		}
		dict[num] = i
	}
	return res
}
