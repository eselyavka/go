package solutions

func takeCharacters(s string, k int) int {
	n := len(s)
	cnt := make([]int, 3, 3)
	for i := 0; i < n; i++ {
		cnt[int(s[i])-int('a')]++
	}

	for _, v := range cnt {
		if v < k {
			return -1
		}
	}

	left := 0
	max_window := 0
	window := make([]int, 3, 3)

	for right := 0; right < n; right++ {
		window[int(s[right])-int('a')]++

		for left <= right && (cnt[0]-window[0] < k || cnt[1]-window[1] < k || cnt[2]-window[2] < k) {
			window[int(s[left])-int('a')]--
			left++
		}

		max_window = MaxInts([]int{max_window, right - left + 1})
	}

	return n - max_window
}
