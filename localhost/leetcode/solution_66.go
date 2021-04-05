package solutions

func plusOne(digits []int) []int {
	n := len(digits)
	var res = make([]int, n)
	carry := 1

	for i := n - 1; i >= 0; i-- {
		sum := (digits[i] + carry) % 10
		carry = (digits[i] + carry) / 10
		res[i] = sum
	}

	if carry == 1 {
		res = append([]int{1}, res...)
	}

	return res
}
