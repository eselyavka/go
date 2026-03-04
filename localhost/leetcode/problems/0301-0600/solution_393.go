package solutions

func validUtf8(data []int) bool {
	if len(data) == 0 {
		return false
	}

	leadingOnesByte := func(x int) int {
		b := x & 0xFF

		if (b & 0b10000000) == 0 {
			return 0
		}
		if (b & 0b11100000) == 0b11000000 {
			return 2
		}
		if (b & 0b11110000) == 0b11100000 {
			return 3
		}
		if (b & 0b11111000) == 0b11110000 {
			return 4
		}
		return -1
	}

	n := len(data)
	i := 0

	for i < n {
		ones := leadingOnesByte(data[i])
		if ones == -1 {
			return false
		}
		if ones == 1 || ones > 4 {
			return false
		}

		if ones == 0 {
			i++
			continue
		}

		if i+ones > n {
			return false
		}

		for j := 1; j < ones; j++ {
			b := data[i+j] & 0xFF
			if (b & 0b11000000) != 0b10000000 {
				return false
			}
		}

		i += ones
	}

	return true
}
