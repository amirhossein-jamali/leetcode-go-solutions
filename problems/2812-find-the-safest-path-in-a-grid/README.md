---
id: 2812
title: "Find the Safest Path in a Grid"
difficulty: "Medium"
level: "Staff"
platform: "LeetCode"
link: "https://leetcode.com/problems/find-the-safest-path-in-a-grid/"
contest: "Weekly Contest 357"
status: "Solved"
language: "Go"
topics:
  - "Array"
  - "Binary Search"
  - "Breadth-First Search"
  - "Union Find"
  - "Matrix"
go_concepts:
  - "Manhattan distance transform on a grid"
  - "Bucketed intrusive linked lists inside the matrix"
  - "Disjoint Set Union with path compression"
  - "Descending safeness activation with connectivity checks"
  - "Table-driven tests with grid cloning"
tags:
  - leetcode
  - go
  - array
  - union-find
  - matrix
  - shortest-path
  - weekly-contest-357
---

# 2812. Find the Safest Path in a Grid

## Problem Link

LeetCode: `https://leetcode.com/problems/find-the-safest-path-in-a-grid/`

## Difficulty

Medium

## Problem Topics

* Array
* Binary Search
* Breadth-First Search
* Union Find
* Matrix

## What to Know Before Solving

General concepts:

* The safeness factor of a path is the minimum Manhattan distance from any cell on the path to any thief
* We want the maximum possible safeness factor from `(0, 0)` to `(n - 1, n - 1)`
* First compute each cell's distance to the nearest thief
* Then activate cells from highest distance downward until start and end become connected
* DSU can track connectivity as cells become available

Go concepts:

* Two-pass Manhattan distance transform on a grid
* Reusing the grid for multiple roles: distances, linked lists, and DSU state
* Fixed-size bucket heads with `[maxLevels]int32`
* DSU helpers with path compression: `findRoot` and `union`
* Cloning `[][]int` in tests because the solution mutates `grid`

## Problem Description

You are given a `0`-indexed `2D` matrix `grid` of size `n x n`, where:

* `grid[r][c] = 1` means the cell contains a thief
* `grid[r][c] = 0` means the cell is empty

You start at `(0, 0)` and want to reach `(n - 1, n - 1)` using 4-directional moves.

The safeness factor of a path is the minimum Manhattan distance from any cell in the path to any thief in the grid.

Return the maximum safeness factor of all paths leading to cell `(n - 1, n - 1)`.

## Function Signature

Expected LeetCode function signature:

```go
func maximumSafenessFactor(grid [][]int) int {

}
```

## Examples

### Example 1

Input:

```text
grid = [[1,0,0],[0,0,0],[0,0,1]]
```

Output:

```text
0
```

Explanation:

```text
All paths from (0, 0) to (n - 1, n - 1) go through thief cells.
```

### Example 2

Input:

```text
grid = [[0,0,1],[0,0,0],[0,0,0]]
```

Output:

```text
2
```

Explanation:

```text
The best path has safeness factor 2.
The closest thief to cell (0, 0) on that path is at (0, 2), with distance 2.
```

### Example 3

Input:

```text
grid = [[0,0,0,1],[0,0,0,0],[0,0,0,0],[1,0,0,0]]
```

Output:

```text
2
```

Explanation:

```text
The best path has safeness factor 2 relative to both corner thieves.
```

## Constraints

```text
1 <= grid.length == n <= 400
grid[i].length == n
grid[i][j] is either 0 or 1.
There is at least one thief in the grid.
```

## Approach

### Step 1: Manhattan distance transform

Convert the grid so each cell stores its Manhattan distance to the nearest thief.

Thief cells become distance `0`. Empty cells start as a large value and are updated with forward and backward passes.

### Step 2: Bucket cells by distance

Group cells into buckets by their distance value using intrusive linked lists stored inside the grid.

### Step 3: Activate cells from high to low safeness

Process safeness levels from largest to smallest.

When a cell becomes active, union it with already active neighbors.

After each level, check whether `(0, 0)` and `(n - 1, n - 1)` are connected.

The first safeness level where they connect is the answer.

## Algorithm

1. If start or end is a thief, return `0`.
2. Build Manhattan distances with two grid passes.
3. Place each cell into a distance bucket using linked lists in `grid`.
4. For `safeness` from `maxDistance` down to `0`:
   * Activate every cell in that bucket.
   * Union with active neighbors.
   * If start and end are in the same DSU component, return `safeness`.
5. Return `0` if no connection is found.

## Why This Works

For any path, its safeness factor is limited by the weakest cell on that path, meaning the cell with smallest distance-to-thief value.

If we only allow cells with distance at least `d`, then a path with safeness factor at least `d` exists if and only if start and end are connected through such cells.

Activating cells from high distance downward finds the largest `d` where connectivity is possible.

Therefore, the first connected safeness level is the maximum safeness factor.

## Complexity Analysis

Let `n` be the grid size and `N = n^2`.

### Time Complexity

```text
O(n^2 * alpha(n^2))
```

Distance transform is `O(n^2)`. Each cell is activated once, and DSU operations are nearly constant.

### Space Complexity

```text
O(n^2)
```

The algorithm reuses the input grid and a fixed bucket-head array.

Note: the solution mutates `grid`.

## Code

The Go solution is available in:

```text
solution.go
```

Important implementation details:

* Start/end thief cells immediately return `0`
* `maxLevels = 799` is enough because `n <= 400`
* Negative values in `grid` encode DSU roots and parent pointers
* `findRoot` uses path compression for efficiency

## Test Cases

The local tests are available in:

```text
solution_test.go
```

Test coverage includes:

* All three LeetCode examples
* A grid where the start cell is a thief

## Edge Cases

Important cases to consider:

* Start or end on a thief
* Only one thief in the grid
* Large grid with distant thieves
* Path must go around thieves to maximize safeness
* Minimum `n = 1` is excluded by constraints, but small `2 x 2` grids matter

## Notes

* This combines distance transform, bucketed traversal, and union-find.
* The in-place encoding is memory-efficient but advanced.
* Copy `grid` in local tests because the function modifies the input matrix.

## Related Problem

* [3286. Find a Safe Walk Through a Grid](./../3286-find-a-safe-walk-through-a-grid/) — another grid path problem with safety constraints, using 0-1 BFS instead.
