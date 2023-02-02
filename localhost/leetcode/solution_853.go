package solutions

import "sort"

func carFleet(target int, position []int, speed []int) int {
	n := len(position)

	pos_and_speed := make([][]int, 0)
	for i := 0; i < n; i++ {
		arr := []int{position[i], speed[i]}
		pos_and_speed = append(pos_and_speed, arr)
	}

	sort.Slice(pos_and_speed, func(i, j int) bool {
		return pos_and_speed[i][0] < pos_and_speed[j][0]
	})

	stack := make([][]int, 0)

	for i := n - 1; i > -1; i-- {
		m := len(stack)
		if m > 0 && float64((target-stack[m-1][0]))/float64(stack[m-1][1]) >= float64((target-pos_and_speed[i][0]))/float64(pos_and_speed[i][1]) {
			continue
		}
		stack = append(stack, pos_and_speed[i])
	}

	return len(stack)
}
