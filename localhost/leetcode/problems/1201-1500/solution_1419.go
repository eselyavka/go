package solutions

func minNumberOfFrogs(croakOfFrogs string) int {
	if len(croakOfFrogs) < 5 {
		return -1
	}

	c := 0
	r := 0
	o := 0
	a := 0
	k := 0

	frogs := 0
	ans := 0

	for _, x := range croakOfFrogs {
		if x == 'c' {
			c++
			frogs++
			ans = MaxInts([]int{ans, frogs})
		} else if x == 'r' {
			r++
		} else if x == 'o' {
			o++
		} else if x == 'a' {
			a++
		} else {
			k++
			frogs--
		}

		if c < r || r < o || o < a || a < k {
			return -1
		}
	}

	if frogs == 0 && c == r && r == o && o == a && a == k {
		return ans
	}

	return -1
}
