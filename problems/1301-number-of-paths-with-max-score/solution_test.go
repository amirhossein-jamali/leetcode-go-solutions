package numberofpathswithmaxscore

import (
	"math/rand"
	"reflect"
	"testing"
)

func referencePathsWithMaxScore(board []string) []int {
	const mod = 1_000_000_007

	n := len(board)
	score := make([][]int, n)
	ways := make([][]int, n)

	for row := range score {
		score[row] = make([]int, n)
		ways[row] = make([]int, n)

		for col := range score[row] {
			score[row][col] = -1
		}
	}

	score[n-1][n-1] = 0
	ways[n-1][n-1] = 1

	for row := n - 1; row >= 0; row-- {
		for col := n - 1; col >= 0; col-- {
			if board[row][col] == 'X' {
				continue
			}

			if row == n-1 && col == n-1 {
				continue
			}

			bestScore := -1
			bestWays := 0

			update := func(candidateScore, candidateWays int) {
				if candidateScore < 0 {
					return
				}

				if candidateScore > bestScore {
					bestScore = candidateScore
					bestWays = candidateWays
					return
				}

				if candidateScore == bestScore {
					bestWays += candidateWays

					if bestWays >= mod {
						bestWays -= mod
					}
				}
			}

			if row+1 < n {
				update(score[row+1][col], ways[row+1][col])
			}

			if col+1 < n {
				update(score[row][col+1], ways[row][col+1])
			}

			if row+1 < n && col+1 < n {
				update(score[row+1][col+1], ways[row+1][col+1])
			}

			if bestScore < 0 {
				continue
			}

			cell := board[row][col]
			if cell != 'E' && cell != 'S' {
				bestScore += int(cell - '0')
			}

			score[row][col] = bestScore
			ways[row][col] = bestWays
		}
	}

	if score[0][0] < 0 {
		return []int{0, 0}
	}

	return []int{score[0][0], ways[0][0]}
}

func TestPathsWithMaxScore(t *testing.T) {
	tests := []struct {
		name  string
		board []string
		want  []int
	}{
		{
			name:  "leetcode example 1",
			board: []string{"E23", "2X2", "12S"},
			want:  []int{7, 1},
		},
		{
			name:  "leetcode example 2",
			board: []string{"E12", "1X1", "21S"},
			want:  []int{4, 2},
		},
		{
			name:  "leetcode example 3",
			board: []string{"E11", "XXX", "11S"},
			want:  []int{0, 0},
		},
		{
			name:  "minimum 2x2 board with two equal optimal paths",
			board: []string{"E1", "1S"},
			want:  []int{1, 2},
		},
		{
			name:  "unique valid path",
			board: []string{"E9", "XS"},
			want:  []int{9, 1},
		},
		{
			name:  "no path because all moves from start are blocked",
			board: []string{"EXX", "XXX", "XXS"},
			want:  []int{0, 0},
		},
		{
			name:  "only maximum-score routes to intermediate cell count",
			board: []string{"E22", "121", "21S"},
			want:  []int{5, 3},
		},
		{
			name:  "locally largest next digit is not globally optimal",
			board: []string{"E11", "9X9", "98S"},
			want:  []int{26, 1},
		},
		{
			name:  "optimal route uses a diagonal move",
			board: []string{"E19", "911", "19S"},
			want:  []int{19, 2},
		},
		{
			name:  "larger open board exercises modular path counting",
			board: []string{"E1111", "11111", "11111", "11111", "1111S"},
			want:  []int{7, 70},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := pathsWithMaxScore(tt.board)
			if !reflect.DeepEqual(got, tt.want) {
				t.Fatalf(
					"%s: pathsWithMaxScore(%v) = %v, want %v",
					tt.name,
					tt.board,
					got,
					tt.want,
				)
			}

			reference := referencePathsWithMaxScore(tt.board)
			if !reflect.DeepEqual(got, reference) {
				t.Fatalf(
					"%s: reference = %v, optimized = %v",
					tt.name,
					reference,
					got,
				)
			}
		})
	}
}

func TestPathsWithMaxScoreRandomized(t *testing.T) {
	rng := rand.New(rand.NewSource(42))

	for trial := 0; trial < 300; trial++ {
		n := 2 + rng.Intn(5)

		board := make([][]byte, n)
		for row := range board {
			board[row] = make([]byte, n)
		}

		board[0][0] = 'E'
		board[n-1][n-1] = 'S'

		for row := 0; row < n; row++ {
			for col := 0; col < n; col++ {
				if row == 0 && col == 0 {
					continue
				}

				if row == n-1 && col == n-1 {
					continue
				}

				switch rng.Intn(11) {
				case 0:
					board[row][col] = 'X'
				default:
					board[row][col] = byte('0' + 1 + rng.Intn(9))
				}
			}
		}

		encoded := make([]string, n)
		for row := range board {
			encoded[row] = string(board[row])
		}

		got := pathsWithMaxScore(encoded)
		want := referencePathsWithMaxScore(encoded)

		if !reflect.DeepEqual(got, want) {
			t.Fatalf(
				"trial %d: pathsWithMaxScore(%v) = %v, want %v",
				trial,
				encoded,
				got,
				want,
			)
		}
	}
}
