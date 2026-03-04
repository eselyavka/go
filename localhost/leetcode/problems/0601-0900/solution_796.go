package solutions

func rec(x string, y string, i int) bool {
	if x == y {
		return true
	}

	if i >= len(x) {
		return false
	}

	return rec(x[1:]+x[0:1], y, i+1)
}

func rotateString(s string, goal string) bool {
	return rec(s, goal, 0)
}
