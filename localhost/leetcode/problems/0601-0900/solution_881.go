package solutions

import (
	"sort"
)

func numRescueBoats(people []int, limit int) int {
	ans := 0

	sort.Ints(people)

	for len(people) > 0 {
		curr := people[len(people)-1]
		people = people[:len(people)-1]

		next := 0
		if len(people) > 0 {
			next = people[0]
			people = people[1:]
		}

		if curr+next > limit {
			people = append([]int{next}, people...)
		}
		ans += 1
	}

	return ans
}
