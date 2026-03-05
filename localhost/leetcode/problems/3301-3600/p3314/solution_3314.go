package p3314

func minBitwiseArray(nums []int) []int {
	ans := make([]int, len(nums))

	for idx, num := range nums {
		if num%2 == 0 {
			ans[idx] = -1
		} else {
			r := 0
			t := num
			for t&1 != 0 {
				r++
				t >>= 1
			}
			ans[idx] = num - (1 << (r - 1))
		}
	}

	return ans
}
