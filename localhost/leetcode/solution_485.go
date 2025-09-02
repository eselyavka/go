package solutions

func findMaxConsecutiveOnes(nums []int) int {
	runningSum := 0
	prevSum := 0
	for _, n := range nums {
		if n == 1 {
			runningSum += 1
		} else {
			prevSum = MaxInts([]int{prevSum, runningSum})
			runningSum = 0
		}
	}

	return MaxInts([]int{runningSum, prevSum})
}
