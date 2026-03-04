package solutions

func rotateTheBox(box [][]byte) [][]byte {
	m := len(box)
	n := len(box[0])
	res := make([][]byte, n)

	for i := 0; i < m; i++ {
		l := 0
		r := 0
		for l < n {
			if box[i][l] == '.' {
				tmp := box[i][l]
				box[i][l] = box[i][r]
				box[i][r] = tmp
				r++
			}
			if box[i][l] == '*' {
				r = l + 1
			}
			l++
		}

		p := 0
		for p < n {
			res[p] = make([]byte, m, m)
			p++
		}

		k := 0
		for i := m - 1; i > -1; i-- {
			for j := 0; j < n; j++ {
				res[j][k] = box[i][j]
			}
			k++
		}
	}

	return res
}
