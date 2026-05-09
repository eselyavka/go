package p752

type stateStep struct {
	state string
	steps int
}

func openLock(deadends []string, target string) int {
	if target == "0000" {
		return 0
	}

	dead := make(map[string]struct{})
	for _, deadend := range deadends {
		dead[deadend] = struct{}{}
	}

	if _, ok := dead["0000"]; ok {
		return -1
	}

	if _, ok := dead[target]; ok {
		return -1
	}

	neighbors := func(state string) []string {
		result := make([]string, 0, 8)
		for i := 0; i < 4; i++ {
			digit := int(state[i] - '0')
			for _, delta := range []int{1, -1} {
				nextDigit := (digit + delta + 10) % 10
				nextState := state[:i] + string(rune('0'+nextDigit)) + state[i+1:]
				result = append(result, nextState)
			}
		}
		return result
	}

	queue := []stateStep{{state: "0000", steps: 0}}
	visited := map[string]struct{}{"0000": {}}

	for len(queue) > 0 {
		item := queue[0]
		queue = queue[1:]

		for _, neighbor := range neighbors(item.state) {
			if _, ok := dead[neighbor]; ok {
				continue
			}
			if _, ok := visited[neighbor]; ok {
				continue
			}
			if neighbor == target {
				return item.steps + 1
			}
			visited[neighbor] = struct{}{}
			queue = append(queue, stateStep{state: neighbor, steps: item.steps + 1})
		}
	}

	return -1
}
