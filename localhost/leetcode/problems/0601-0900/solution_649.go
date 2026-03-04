package solutions

func predictPartyVictory(senate string) string {
	n := len(senate)

	R := make([]int, 0)
	D := make([]int, 0)

	for idx, c := range senate {
		if c == 'R' {
			R = append(R, idx)
		} else {
			D = append(D, idx)
		}
	}

	for len(R) > 0 && len(D) > 0 {
		r_idx := R[0]
		R = R[1:]
		d_idx := D[0]
		D = D[1:]

		if r_idx < d_idx {
			R = append(R, r_idx+n)
		} else {
			D = append(D, d_idx+n)
		}
	}

	if len(R) > 0 {
		return "Radiant"
	}

	return "Dire"
}
