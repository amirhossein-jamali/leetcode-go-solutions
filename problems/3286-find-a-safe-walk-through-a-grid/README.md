---
id: 3286
title: "Find a Safe Walk Through a Grid"
difficulty: "Medium"
level: "Staff"
platform: "LeetCode"
link: "https://leetcode.com/problems/find-a-safe-walk-through-a-grid/"
contest: "Biweekly Contest 139"
status: "Solved"
language: "Go"
topics:
  - "Array"
  - "Breadth-First Search"
  - "Graph Theory"
  - "Matrix"
  - "Shortest Path"
go_concepts:
  - "0-1 BFS with two linked queues"
  - "In-place grid marking for visited cells"
  - "Linear index encoding for matrix cells"
  - "Direction array for four neighbors"
  - "Table-driven tests with grid cloning"
tags:
  - leetcode
  - go
  - array
  - bfs
  - matrix
  - shortest-path
  - biweekly-contest-139
---

# 3286. Find a Safe Walk Through a Grid

## Problem Link

LeetCode: `https://leetcode.com/problems/find-a-safe-walk-through-a-grid/`

## Difficulty

Medium

## Problem Topics

* Array
* Breadth-First Search
* Graph Theory
* Matrix
* Shortest Path

## What to Know Before Solving

General concepts:

* Each unsafe cell (`1`) costs `1` health point
* You must reach the bottom-right cell with health at least `1`
* This is a shortest-path problem on a grid with two edge costs: `0` and `1`
* A 0-1 BFS processes all zero-cost moves before one-cost moves
* Because health is limited, unsafe cells can only be entered when enough health remains

Go concepts:

* Encoding `(row, col)` as a single index: `row*cols + col`
* Using the grid itself to store visited state and linked-list pointers
* Two-queue style traversal with `currentHead` and `nextHead`
* A compact direction array: `{-1, 0, 1, 0, -1}`
* Cloning `[][]int` in tests because the solution mutates `grid`

## Problem Description

You are given an `m x n` binary matrix `grid` and an integer `health`.

You start on the upper-left corner `(0, 0)` and would like to get to the lower-right corner `(m - 1, n - 1)`.

You can move up, down, left, or right from one cell to another adjacent cell as long as your health remains positive.

Cells `(i, j)` with `grid[i][j] = 1` are considered unsafe and reduce your health by `1`.

Return `true` if you can reach the final cell with a health value of `1` or more, and `false` otherwise.

## Function Signature

Expected LeetCode function signature:

```go
func findSafeWalk(grid [][]int, health int) bool {

}
```

## Examples

### Example 1

Input:

```text
grid = [[0,1,0,0,0],[0,1,0,1,0],[0,0,0,1,0]], health = 1
```

Output:

```text
true
```

Explanation:

```text
The final cell can be reached safely by walking along safe cells.
```

### Example 2

Input:

```text
grid = [[0,1,1,0,0,0],[1,0,1,0,0,0],[0,1,1,1,0,1],[0,0,1,0,1,0]], health = 3
```

Output:

```text
false
```

Explanation:

```text
A minimum of 4 health points is needed to reach the final cell safely.
```

### Example 3

Input:

```text
grid = [[1,1,1],[1,0,1],[1,1,1]], health = 5
```

Output:

```text
true
```

Explanation:

```text
The final cell can be reached safely by going through the center safe cell (1, 1).
Any path that avoids that cell is unsafe.
```

## Constraints

```text
m == grid.length
n == grid[i].length
1 <= m, n <= 50
2 <= m * n
1 <= health <= m + n
grid[i][j] is either 0 or 1.
```

## Approach

Treat the grid as a graph:

* Safe cells (`0`) have cost `0`
* Unsafe cells (`1`) have cost `1`

We want the path with minimum total unsafe-cell cost from start to end.

Use 0-1 BFS:

* Process all reachable safe cells first at the current health level
* Then decrease health by `1` and process cells that require entering one unsafe cell

Instead of a normal queue and visited matrix, the solution stores:

* visited marks as negative values
* linked-list pointers inside the grid to chain the next cell in the current BFS layer

## Algorithm

1. Subtract the starting cell cost from `health`.
2. If health is already `0` or less, return `false`.
3. Mark the start cell visited and begin 0-1 BFS with `currentHead` and `nextHead`.
4. For each cell dequeued:
   * If it is the target, return `true`.
   * Try all four neighbors.
   * Skip out-of-bounds, visited, or unsafe cells when `health <= 1`.
   * Safe neighbors join the current zero-cost layer.
   * Unsafe neighbors join the next one-cost layer.
5. When the current layer is empty, spend `1` health and move to the next layer.
6. If all layers are exhausted, return `false`.

## Why This Works

Among all paths, the one that minimizes unsafe-cell visits gives the best chance to survive.

0-1 BFS explores cells in nondecreasing path cost order, so the first time we can reach the destination, we have used the minimum possible number of unsafe cells.

The health check `cell == 1 && health <= 1` ensures we never enter an unsafe cell unless we can still finish with positive health.

Therefore, returning `true` means a safe walk exists, and returning `false` means none exists.

## Complexity Analysis

Let `m` be the number of rows and `n` be the number of columns.

### Time Complexity

```text
O(m * n * health)
```

Each cell can be processed across multiple health layers. With the given constraints, this is efficient enough.

### Space Complexity

```text
O(1)
```

The traversal reuses the input grid for visited marks and queue links.

Note: the solution mutates `grid`.

## Code

The Go solution is available in:

```text
solution.go
```

Important implementation details:

* Cells are encoded as `row*cols + col`
* Negative grid values store both visited state and next-pointer data
* Safe and unsafe neighbors are appended to different BFS layers
* The target cell is checked both when dequeuing and when enqueueing neighbors

## Test Cases

The local tests are available in:

```text
solution_test.go
```

Test coverage includes:

* All three LeetCode examples
* Failing immediately at the starting cell
* A grid with only safe cells

## Edge Cases

Important cases to consider:

* Starting cell is unsafe and consumes all health
* Target reached through only safe cells
* Need to pass through multiple unsafe cells
* Not enough health to enter required unsafe cells
* Small `2 x 2` grids

## Notes

* This is a 0-1 BFS problem disguised as a grid walk with health.
* The in-place linked-list queue is memory-efficient but harder to read than a standard deque.
* Copy `grid` in local tests because the function modifies the input matrix.
