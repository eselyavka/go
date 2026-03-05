package solutions

func char_arr_cmp(arr1 []int, arr2 []int) bool {
	for i := 0; i < 26; i++ {
		if arr1[i] != arr2[i] {
			return false
		}
	}

	return true
}
func checkInclusion(s1 string, s2 string) bool {
	n := len(s1)
	m := len(s2)

	if n > m {
		return false
	}

	s1_arr := make([]int, 26, 26)
	for _, c := range s1 {
		s1_arr[int(c)-int('a')]++
	}

	for i := 0; i <= m-n; i++ {
		s2_arr := make([]int, 26, 26)
		for j := 0; j < n; j++ {
			s2_arr[int(s2[i+j])-int('a')]++
		}
		if char_arr_cmp(s1_arr, s2_arr) {
			return true
		}
	}
	return false
}
