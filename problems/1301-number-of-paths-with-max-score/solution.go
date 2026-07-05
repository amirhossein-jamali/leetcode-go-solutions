package numberofpathswithmaxscore

func pathsWithMaxScore(board []string) []int {
	const mod uint32 = 1_000_000_007

	n := len(board)

	var score [100]uint16
	var ways [100]uint32

	score[n-1] = 1
	ways[n-1] = 1

	for row := n - 1; row >= 0; row-- {
		var rightScore uint16
		var rightWays uint32
		var diagonalScore uint16
		var diagonalWays uint32

		startCol := n - 1

		if row == n-1 {
			rightScore = 1
			rightWays = 1
			startCol = n - 2
		}

		for col := startCol; col >= 0; col-- {
			belowScore := score[col]
			belowWays := ways[col]
			cell := board[row][col]

			if cell == 'X' {
				score[col] = 0
				ways[col] = 0

				rightScore = 0
				rightWays = 0

				diagonalScore = belowScore
				diagonalWays = belowWays
				continue
			}

			bestScore := belowScore
			bestWays := belowWays

			if rightScore > bestScore {
				bestScore = rightScore
				bestWays = rightWays
			} else if rightScore == bestScore {
				bestWays += rightWays

				if bestWays >= mod {
					bestWays -= mod
				}
			}

			if diagonalScore > bestScore {
				bestScore = diagonalScore
				bestWays = diagonalWays
			} else if diagonalScore == bestScore {
				bestWays += diagonalWays

				if bestWays >= mod {
					bestWays -= mod
				}
			}

			if bestScore == 0 {
				score[col] = 0
				ways[col] = 0

				rightScore = 0
				rightWays = 0
			} else {
				if cell != 'E' {
					bestScore += uint16(cell - '0')
				}

				score[col] = bestScore
				ways[col] = bestWays

				rightScore = bestScore
				rightWays = bestWays
			}

			diagonalScore = belowScore
			diagonalWays = belowWays
		}
	}

	if score[0] == 0 {
		return []int{0, 0}
	}

	return []int{
		int(score[0] - 1),
		int(ways[0]),
	}
}
