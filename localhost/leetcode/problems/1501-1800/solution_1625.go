package solutions

func findLexSmallestString(s string, a, b int) string {
	n := len(s)

	rotate := func(x string) string {
		return x[n-b:] + x[:n-b]
	}

	addOdd := func(x string) string {
		bytes := []byte(x)
		for i := 0; i < n; i++ {
			if i%2 == 1 {
				d := int(bytes[i] - '0')
				d = (d + a) % 10
				bytes[i] = byte('0' + d)
			}
		}
		return string(bytes)
	}

	seen := make(map[string]bool)
	stack := []string{s}

	best := ""

	minStr := func(a, b string) string {
		if a == "" || b < a {
			return b
		}
		return a
	}

	for len(stack) > 0 {
		cur := stack[len(stack)-1]
		stack = stack[:len(stack)-1]

		if seen[cur] {
			continue
		}
		seen[cur] = true
		best = minStr(best, cur)

		r := rotate(cur)
		ad := addOdd(cur)
		if !seen[r] {
			stack = append(stack, r)
		}
		if !seen[ad] {
			stack = append(stack, ad)
		}
	}

	return best
}
