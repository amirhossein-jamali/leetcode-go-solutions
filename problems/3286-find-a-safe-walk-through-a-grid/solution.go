package findasafewalkthroughagrid

func findSafeWalk(grid [][]int, health int) bool {
	rows := len(grid)
	cols := len(grid[0])
	target := rows*cols - 1

	health -= grid[0][0]
	if health <= 0 {
		return false
	}

	currentHead := 0
	nextHead := -1

	grid[0][0] = -1

	directions := [5]int{-1, 0, 1, 0, -1}

	for currentHead != -1 || nextHead != -1 {
		if currentHead == -1 {
			health--
			if health <= 0 {
				return false
			}

			currentHead = nextHead
			nextHead = -1
		}

		index := currentHead
		row := index / cols
		col := index % cols

		currentHead = -grid[row][col] - 2

		if index == target {
			return true
		}

		for direction := 0; direction < 4; direction++ {
			nextRow := row + directions[direction]
			nextCol := col + directions[direction+1]

			if nextRow < 0 || nextRow >= rows ||
				nextCol < 0 || nextCol >= cols {
				continue
			}

			cell := grid[nextRow][nextCol]

			if cell < 0 {
				continue
			}

			if cell == 1 && health <= 1 {
				continue
			}

			nextIndex := nextRow*cols + nextCol

			if nextIndex == target {
				return true
			}

			if cell == 0 {
				grid[nextRow][nextCol] = -(currentHead + 2)
				currentHead = nextIndex
			} else {
				grid[nextRow][nextCol] = -(nextHead + 2)
				nextHead = nextIndex
			}
		}
	}

	return false
}
