package p874

import "github.com/eseliavka/go/util"

func robotSim(commands []int, obstacles [][]int) int {
	obstacleSet := make(map[util.TupleInt]struct{}, len(obstacles))
	for _, obstacle := range obstacles {
		obstacleSet[util.TupleInt{Row: obstacle[0], Col: obstacle[1]}] = struct{}{}
	}

	directions := [][2]int{{0, 1}, {1, 0}, {0, -1}, {-1, 0}}
	direction := 0
	x, y := 0, 0
	best := 0

	for _, command := range commands {
		switch command {
		case -1:
			direction = (direction + 1) % len(directions)
		case -2:
			direction = (direction + len(directions) - 1) % len(directions)
		default:
			for step := 0; step < command; step++ {
				nextX := x + directions[direction][0]
				nextY := y + directions[direction][1]
				if _, blocked := obstacleSet[util.TupleInt{Row: nextX, Col: nextY}]; blocked {
					break
				}

				x = nextX
				y = nextY
				distance := x*x + y*y
				if distance > best {
					best = distance
				}
			}
		}
	}

	return best
}
