package solutions

func checkDistances(s string, distance []int) bool {
	d := make(map[int]int)

	for i, c := range s {
		idx := int(c) - int('a')

		if val, ok := d[idx]; ok {
			d[idx] = i - val - 1
		} else {
			d[idx] = i
		}
	}

	for k, v := range d {
		if distance[k] != v {
			return false
		}
	}

	return true
}
