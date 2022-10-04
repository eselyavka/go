package solutions

func longestConsecutive(nums []int) int {
	set := map[int]struct{}{}

	for _, num := range nums {
		set[num] = struct{}{}
	}

	longest_len := 0

	i := 0
	for i < len(nums) {
		if _, ok := set[nums[i]-1]; !ok {
			curr_len := 1
			next_num := nums[i] + 1
			_, ok := set[next_num]
			for ok == true {
				curr_len++
				next_num++
				_, ok = set[next_num]
			}

			if curr_len > longest_len {
				longest_len = curr_len
			}
		}
		i++
	}
	return longest_len
}
