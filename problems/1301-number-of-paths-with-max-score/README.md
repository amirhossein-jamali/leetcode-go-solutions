---
id: 1301
title: "Number of Paths with Max Score"
difficulty: "Hard"
level: "Principal"
platform: "LeetCode"
link: "https://leetcode.com/problems/number-of-paths-with-max-score/"
contest: "Biweekly Contest 16"
status: "Solved"
language: "Go"
topics:
  - "Array"
  - "Dynamic Programming"
  - "Matrix"
go_concepts:
  - "Bottom-up dynamic programming on a DAG"
  - "Implicit topological ordering"
  - "One-dimensional rolling DP"
  - "Fixed-size arrays derived from constraints"
  - "Unsigned integer types"
  - "Modulo arithmetic without division"
  - "Table-driven tests"
tags:
  - leetcode
  - go
  - array
  - dynamic-programming
  - matrix
  - dag
  - rolling-dp
  - memory-optimization
  - biweekly-contest-16
---

# 1301. Number of Paths with Max Score

## Problem Link

LeetCode: `https://leetcode.com/problems/number-of-paths-with-max-score/`

## Difficulty

Hard

## Problem Topics

* Array
* Dynamic Programming
* Matrix

## What to Know Before Solving

General concepts:

* Directed acyclic graphs and why move constraints create a DAG on a grid
* Bottom-up dynamic programming with two coupled states: maximum score and path count
* Combining predecessor states by maximum score, then summing only matching counts
* Modulo arithmetic for large path counts
* Why greedy digit selection fails when future path totals matter

Go concepts:

* Bottom-up DP with implicit topological order instead of an explicit graph
* One-dimensional rolling DP over columns within each row
* Fixed-size arrays `[100]uint16` and `[100]uint32` derived from constraints
* Unsigned integer types for compact storage and safe modular addition
* Table-driven tests with a reference two-dimensional DP implementation

## Problem Description

You are given a square board of characters. You start at the bottom-right cell marked `S` and must reach the top-left cell marked `E`.

Each other cell contains either:

* a digit from `'1'` to `'9'`, representing its score
* or `'X'`, representing an obstacle

From a cell, you may move only:

* up
* left
* diagonally up-left

You may not enter an obstacle.

Return two integers:

1. The maximum score obtainable along a valid path from `S` to `E`.
2. The number of paths that obtain exactly that maximum score, modulo `1_000_000_007`.

The `S` and `E` cells contribute zero score.

If no valid path exists, return `[0, 0]`.

## Function Signature

Expected LeetCode function signature:

```go
func pathsWithMaxScore(board []string) []int {

}
```

## Examples

### Example 1

Input:

```text
board = ["E23","2X2","12S"]
```

Output:

```text
[7,1]
```

Explanation:

```text
The only maximum-score path is:

S -> (1,2) -> (0,2) -> (0,1) -> E

Score: 2 + 3 + 2 = 7
```

### Example 2

Input:

```text
board = ["E12","1X1","21S"]
```

Output:

```text
[4,2]
```

Explanation:

```text
Two maximum-score paths exist, each with score 4:

S -> (2,1) -> (1,0) -> E
S -> (1,2) -> (0,1) -> E
```

### Example 3

Input:

```text
board = ["E11","XXX","11S"]
```

Output:

```text
[0,0]
```

Explanation:

```text
The middle row is entirely blocked, so no valid path exists.
```

## Constraints

```text
2 <= board.length == board[i].length <= 100
```

## Approach

Treat the board as a DAG. Every allowed move strictly decreases `r + c`, so cycles are impossible.

Process cells in reverse topological order: rows from `n - 1` down to `0`, columns from `n - 1` down to `0`. When a cell is processed, all three possible predecessors are already finalized:

* below: `(row + 1, col)`
* right: `(row, col + 1)`
* diagonal: `(row + 1, col + 1)`

For each cell, take the maximum predecessor score, sum path counts only from predecessors tied for that maximum, then add the current digit unless the cell is `S` or `E`.

Compress the two-dimensional tables into one rolling column frontier per row.

## Algorithm

1. Initialize the start cell `S` at `(n - 1, n - 1)` as reachable with encoded score `1` and `1` path.
2. Iterate `row` from `n - 1` down to `0`.
3. Maintain rolling right and diagonal predecessor states for the current row.
4. Iterate `col` from `n - 1` down to `0` (skipping `S` on the first pass of the bottom row).
5. If the cell is `'X'`, mark it unreachable and reset the rolling right state.
6. Otherwise, combine below, right, and diagonal predecessor scores and counts.
7. If no predecessor is reachable, mark the cell unreachable.
8. Otherwise add the digit value for non-`E` cells, store the result, and advance rolling states.
9. After processing `(0, 0)`, return `[score - 1, ways]` or `[0, 0]` if unreachable.

## Why This Works

### DAG interpretation

Treat every non-obstacle cell as a vertex. The possible directed moves are:

```text
(r, c) -> (r - 1, c)
(r, c) -> (r, c - 1)
(r, c) -> (r - 1, c - 1)
```

For every move, `r + c` strictly decreases, so a directed cycle is impossible. Therefore, the board represents a DAG.

### Implicit topological order

Processing rows from `n - 1` down to `0` and columns from `n - 1` down to `0` is a valid topological order.

When processing cell `(row, col)`, all possible predecessor states are already finalized:

* below: `(row + 1, col)`
* right: `(row, col + 1)`
* diagonal: `(row + 1, col + 1)`

No explicit graph or topological-sort data structure is needed.

### DP state

For every cell, conceptually maintain:

* maximum score obtainable from `S` to the cell
* number of paths obtaining that exact maximum score

For each cell:

1. Find the maximum score among below, right, and diagonal.
2. Sum the path counts only from states whose score equals that maximum.
3. Add the current digit to the score.
4. Do not add a value for `S` or `E`.
5. An `X` cell becomes unreachable.

### Score encoding

The implementation stores:

```text
storedScore = actualScore + 1
```

Therefore:

* stored score `0` means unreachable
* stored score `1` means reachable with actual score `0`

This removes the need for signed sentinel values or an initialization loop.

### Rolling one-dimensional DP

During each inner iteration over columns:

* `score[col]` and `ways[col]`, before overwrite, represent the finalized cell below `(row + 1, col)`
* `rightScore` and `rightWays` represent the finalized right neighbor `(row, col + 1)` on the current row
* `diagonalScore` and `diagonalWays` represent the finalized diagonal neighbor `(row + 1, col + 1)`

After processing `(row, col)`, the rolling right state becomes the just-computed cell, and the rolling diagonal state becomes the old below state.

Only one column-sized frontier is required.

### Numeric type safety

Maximum number of numeric cells on a path:

```text
2n - 3
```

For `n <= 100`:

```text
2(100) - 3 = 197
```

Maximum possible score:

```text
197 * 9 = 1773
```

The encoded maximum is `1774`, so `uint16` is safe.

Path counts are stored modulo `1_000_000_007`.

During one modular addition, both operands are below the modulus, so the temporary sum is below:

```text
2 * 1_000_000_007
```

This is below the maximum `uint32` value. Therefore `uint32` is safe.

### Modulo optimization

After each addition:

```go
ways += candidateWays
if ways >= mod {
    ways -= mod
}
```

Both operands are already reduced modulo `mod`, so the sum is at most `2*mod - 2`. Subtracting once when `ways >= mod` is sufficient and avoids the division used by `%`.

## Complexity Analysis

Let:

```text
n = board.length
```

### Time Complexity

```text
Theta(n^2)
```

Every board cell is processed exactly once, and `Theta(n^2)` is optimal because the input may require inspecting every cell.

### Space Complexity

```text
O(n)
```

Conceptually, only one rolling column frontier is needed.

This implementation uses fixed arrays of length `100` because `n <= 100`:

* `[100]uint16` uses `200` raw bytes
* `[100]uint32` uses `400` raw bytes
* total raw DP storage is approximately `600` bytes

Go runtime memory reported by LeetCode includes runtime, stack, allocator, and test-harness overhead and is not equal to only the algorithm's raw DP storage.

## Code

The Go solution is available in:

```text
solution.go
```

Important implementation details:

* `storedScore = actualScore + 1`, so `0` means unreachable
* Rolling `rightScore` / `diagonalScore` track the current row's finalized neighbors
* `uint16` and `uint32` are chosen from constraint bounds, not convenience
* Modular addition uses conditional subtraction instead of `%`
* Fixed local arrays avoid `make`-based slice allocation and may be kept on the stack depending on compiler escape analysis

## Test Cases

The local tests are available in:

```text
solution_test.go
```

Test coverage includes:

* All three LeetCode examples
* Minimum `2 x 2` board with tied optimal paths
* Unique valid path with obstacles
* Fully blocked start with no outgoing moves
* Intermediate-cell merging where only maximum-score routes count
* Board where one-step digit greediness is insufficient
* Optimal route that uses a diagonal move
* Larger open board with many maximum-score paths
* Randomized differential tests against a two-dimensional reference DP

## Edge Cases

Important cases to consider:

* No valid path because obstacles isolate `S` from `E`
* All three moves from `S` blocked immediately
* Single optimal path versus many tied optimal paths
* Cells reachable by multiple routes with different scores
* Boards where the locally largest adjacent digit does not belong to a globally optimal path
* Diagonal moves on the optimal route
* Large path counts requiring modular reduction

## Notes

* A full two-dimensional DP would use `O(n^2)` auxiliary memory.
* A general graph representation and explicit topological sort would add unnecessary memory and overhead.
* Greedy selection of the locally largest digit is not correct because future path scores matter.
* Maximum flow is not the relevant problem because this task maximizes the sum along one path rather than total transferable flow.
* The board is not mutated.
* The fixed `[100]` arrays deliberately use the documented constraint.
