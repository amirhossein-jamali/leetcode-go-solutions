package findthesafestpathinagrid

func maximumSafenessFactor(grid [][]int) int {
	n := len(grid)

	if grid[0][0] == 1 || grid[n-1][n-1] == 1 {
		return 0
	}

	total := n * n
	inf := 2 * n

	// Initialize the Manhattan distance transform.
	for row := 0; row < n; row++ {
		for col := 0; col < n; col++ {
			if grid[row][col] == 1 {
				grid[row][col] = 0
			} else {
				grid[row][col] = inf
			}
		}
	}

	// Forward distance-transform pass.
	for row := 0; row < n; row++ {
		for col := 0; col < n; col++ {
			if row > 0 {
				candidate := grid[row-1][col] + 1
				if candidate < grid[row][col] {
					grid[row][col] = candidate
				}
			}

			if col > 0 {
				candidate := grid[row][col-1] + 1
				if candidate < grid[row][col] {
					grid[row][col] = candidate
				}
			}
		}
	}

	// Backward distance-transform pass.
	for row := n - 1; row >= 0; row-- {
		for col := n - 1; col >= 0; col-- {
			if row+1 < n {
				candidate := grid[row+1][col] + 1
				if candidate < grid[row][col] {
					grid[row][col] = candidate
				}
			}

			if col+1 < n {
				candidate := grid[row][col+1] + 1
				if candidate < grid[row][col] {
					grid[row][col] = candidate
				}
			}
		}
	}

	// n <= 400, therefore the maximum distance is:
	// 2 * (400 - 1) = 798.
	const maxLevels = 799

	var heads [maxLevels]int32
	maxDistance := 2 * (n - 1)

	for distance := 0; distance <= maxDistance; distance++ {
		heads[distance] = -1
	}

	/*
		Build intrusive linked lists inside grid.

		For an inactive cell:

		    grid[row][col] = nextID + 1

		A value of zero means that this is the last node
		in the corresponding bucket.
	*/
	for row := 0; row < n; row++ {
		for col := 0; col < n; col++ {
			distance := grid[row][col]
			id := row*n + col

			grid[row][col] = int(heads[distance]) + 1
			heads[distance] = int32(id)
		}
	}

	target := total - 1

	for safeness := maxDistance; safeness >= 0; safeness-- {
		for id := int(heads[safeness]); id != -1; {
			row := id / n
			col := id % n

			next := grid[row][col] - 1

			// Activate as a DSU root with component size one.
			grid[row][col] = -1

			if row > 0 && grid[row-1][col] < 0 {
				union(grid, n, total, id, id-n)
			}

			if row+1 < n && grid[row+1][col] < 0 {
				union(grid, n, total, id, id+n)
			}

			if col > 0 && grid[row][col-1] < 0 {
				union(grid, n, total, id, id-1)
			}

			if col+1 < n && grid[row][col+1] < 0 {
				union(grid, n, total, id, id+1)
			}

			id = next
		}

		if grid[0][0] < 0 &&
			grid[n-1][n-1] < 0 &&
			findRoot(grid, n, total, 0) ==
				findRoot(grid, n, total, target) {
			return safeness
		}
	}

	return 0
}

func findRoot(grid [][]int, n, total, id int) int {
	root := id

	// Values smaller than -total encode parent IDs.
	for grid[root/n][root%n] < -total {
		root = -grid[root/n][root%n] - total - 1
	}

	// Path compression.
	for id != root {
		row := id / n
		col := id % n

		parent := -grid[row][col] - total - 1
		grid[row][col] = -(total + root + 1)

		id = parent
	}

	return root
}

func union(grid [][]int, n, total, a, b int) {
	rootA := findRoot(grid, n, total, a)
	rootB := findRoot(grid, n, total, b)

	if rootA == rootB {
		return
	}

	rowA, colA := rootA/n, rootA%n
	rowB, colB := rootB/n, rootB%n

	// Root values contain negative component sizes.
	if grid[rowA][colA] > grid[rowB][colB] {
		rootA, rootB = rootB, rootA
		rowA, rowB = rowB, rowA
		colA, colB = colB, colA
	}

	grid[rowA][colA] += grid[rowB][colB]

	// Encode rootA as rootB's parent.
	grid[rowB][colB] = -(total + rootA + 1)
}
