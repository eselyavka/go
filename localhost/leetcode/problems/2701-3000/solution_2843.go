package solutions

import "strconv"

func countSymmetricIntegers(low, high int) int {
	ans := 0

	for i := low; i <= high; i++ {
		str := strconv.Itoa(i)

		if len(str)%2 == 0 {
			half := len(str) / 2
			leftSum, rightSum := 0, 0

			for j := 0; j < half; j++ {
				leftDigit := int(str[j] - '0')
				rightDigit := int(str[j+half] - '0')
				leftSum += leftDigit
				rightSum += rightDigit
			}

			if leftSum == rightSum {
				ans++
			}
		}
	}

	return ans
}
