package solutions

func processGarbage(garbage string) []int {
	res := make([]int, 3, 3)

	for _, c := range garbage {
		if string(c) == "P" {
			res[0]++
		} else if string(c) == "G" {
			res[1]++
		} else {
			res[2]++
		}
	}
	return res
}
func garbageCollection(garbage []string, travel []int) int {
	glass_track := 0
	metal_track := 0
	paper_track := 0

	prefix_travel := make([]int, len(travel), len(travel))
	prefix_travel[0] = travel[0]
	for i := 1; i < len(travel); i++ {
		prefix_travel[i] = prefix_travel[i-1] + travel[i]
	}

	n := len(garbage)

	var metal_idx = -1
	var glass_idx = -1
	var paper_idx = -1

	for i := 0; i < n; i++ {
		process_garbage := processGarbage(garbage[i])
		if process_garbage[0] > 0 {
			paper_track += process_garbage[0]
			if i > 0 {
				paper_idx = i - 1
			}
		}
		if process_garbage[1] > 0 {
			glass_track += process_garbage[1]
			if i > 0 {
				glass_idx = i - 1
			}
		}
		if process_garbage[2] > 0 {
			metal_track += process_garbage[2]
			if i > 0 {
				metal_idx = i - 1
			}
		}
	}
	ans := glass_track + metal_track + paper_track
	if metal_idx != -1 {
		ans += prefix_travel[metal_idx]
	}
	if glass_idx != -1 {
		ans += prefix_travel[glass_idx]
	}
	if paper_idx != -1 {
		ans += prefix_travel[paper_idx]
	}

	return ans
}
