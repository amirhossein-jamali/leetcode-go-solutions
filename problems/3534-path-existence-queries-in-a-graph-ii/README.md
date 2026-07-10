---
id: 3534
title: "Path Existence Queries in a Graph II"
difficulty: "Hard"
level: "Principal"
platform: "LeetCode"
link: "https://leetcode.com/problems/path-existence-queries-in-a-graph-ii/"
contest: "Weekly Contest 447"
status: "Solved"
language: "Go"
topics:
  - "Array"
  - "Two Pointers"
  - "Binary Search"
  - "Dynamic Programming"
  - "Greedy"
  - "Bit Manipulation"
  - "Graph Theory"
  - "Sorting"
go_concepts:
  - "Distinct value compression"
  - "Two-pointer jump table on sorted values"
  - "Binary lifting for shortest path on a value DAG"
  - "Special-case handling for fully connected graphs"
  - "Table-driven tests with BFS reference"
tags:
  - leetcode
  - go
  - array
  - graph
  - binary-lifting
  - two-pointers
  - weekly-contest-447
---

# 3534. Path Existence Queries in a Graph II

## Problem Link

LeetCode: `https://leetcode.com/problems/path-existence-queries-in-a-graph-ii/`

## Difficulty

Hard

## Problem Topics

* Array
* Two Pointers
* Binary Search
* Dynamic Programming
* Greedy
* Bit Manipulation
* Graph Theory
* Sorting

## What to Know Before Solving

General concepts:

* Shortest path in an unweighted graph
* How sorted distinct values form a jump graph
* Two-pointer preprocessing on sorted unique values
* Binary lifting to answer many shortest-path queries
* Special cases when the whole graph is connected or `maxDiff == 0`

Go concepts:

* Value compression with `rankByValue`
* Building a jump table with `uint32`
* Binary lifting over precomputed levels
* BFS reference tests for randomized validation

## Problem Description

You are given `n` nodes labeled `0` to `n - 1`, an array `nums`, and an integer `maxDiff`.

An undirected edge exists between nodes `i` and `j` when:

```text
|nums[i] - nums[j]| <= maxDiff
```

For each query `[ui, vi]`, return the minimum distance between the two nodes. If no path exists, return `-1`.

This is the shortest-path version of [3532. Path Existence Queries in a Graph I](./../3532-path-existence-queries-in-a-graph-i/).

## Function Signature

Expected LeetCode function signature:

```go
func pathExistenceQueries(n int, nums []int, maxDiff int, queries [][]int) []int {

}
```

## Examples

### Example 1

Input:

```text
n = 5, nums = [1,8,3,4,2], maxDiff = 3, queries = [[0,3],[2,4]]
```

Output:

```text
[1,1]
```

### Example 2

Input:

```text
n = 5, nums = [5,3,1,9,10], maxDiff = 2, queries = [[0,1],[0,2],[2,3],[4,3]]
```

Output:

```text
[1,2,-1,1]
```

### Example 3

Input:

```text
n = 3, nums = [3,6,1], maxDiff = 1, queries = [[0,0],[0,1],[1,2]]
```

Output:

```text
[0,-1,-1]
```

## Constraints

```text
1 <= n == nums.length <= 10^5
0 <= nums[i] <= 10^5
0 <= maxDiff <= 10^5
1 <= queries.length <= 10^5
queries[i] == [ui, vi]
0 <= ui, vi < n
```

## Approach

Handle easy special cases first:

* If every value is within `maxDiff` of each other, all nodes are in one connected component and every non-self distance is `1`.
* If `maxDiff == 0`, nodes connect only when their values are equal.

Otherwise:

1. Compress distinct values from `nums` into sorted ranks.
2. Build a jump table on sorted distinct values: from each value, jump to the farthest value still within `maxDiff`.
3. Use binary lifting on that jump table to move toward the target value in logarithmic steps.
4. Convert the lifted distance on values into the shortest node distance answer.

## Algorithm

1. Find `minValue` and `maxValue`.
2. If `maxValue - minValue <= maxDiff`, answer all queries directly.
3. If `maxDiff == 0`, answer using equality of node values.
4. Build sorted distinct `values` and `rankByValue`.
5. Two-pointer preprocess `jump[0][i]` for each distinct value index.
6. Build higher binary-lifting levels.
7. For each query:
   * Handle self, equal values, and direct-edge cases.
   * Lift from the smaller value rank toward the larger one.
   * If reachable, answer is `distance + 1`; otherwise `-1`.

## Why This Works

### Value graph

Only distinct values matter for connectivity. After sorting distinct values `v0 < v1 < ...`, an edge exists between consecutive usable values when their difference is at most `maxDiff`.

The two-pointer jump table stores, for each value rank, the farthest rank reachable in one step while staying within the query range.

### Binary lifting

Once one-step jumps are known, binary lifting precomputes `2^k`-step jumps. Query processing greedily takes the largest possible jump that does not overshoot the target rank.

This gives the minimum number of value jumps between two distinct values.

### Node distance

If two nodes share the same value, distance is `1` when they are different nodes.

If their values differ but are within `maxDiff`, distance is also `1`.

Otherwise the answer comes from the value-jump distance plus one final edge to the target value.

## Complexity Analysis

Let:

```text
n = len(nums)
q = len(queries)
d = number of distinct values in nums
L = number of binary-lifting levels
```

### Time Complexity

```text
O((n + d log d) + q log d)
```

Value compression and jump preprocessing are linear in the value range covered by distinct values. Binary lifting adds a `log d` factor per query.

### Space Complexity

```text
O(d log d)
```

For the jump table and distinct-value arrays.

## Code

The Go solution is available in:

```text
solution.go
```

Important implementation details:

* Early exits avoid building the jump table when unnecessary
* `rankByValue` maps each present value to its compressed rank
* Query answers use value ranks, not raw node indices
* Unreachable targets return `-1`

## Test Cases

The local tests are available in:

```text
solution_test.go
```

Test coverage includes:

* All three LeetCode examples
* Fully connected value-range case
* `maxDiff == 0` with equal values
* Single-node input
* Randomized differential tests against BFS

## Edge Cases

Important cases to consider:

* `u == v`
* Nodes with equal values but different indices
* Values directly within `maxDiff`
* No path between components
* Fully connected graph via value range

## Notes

* Part I only checks connectivity; Part II requires shortest-path distances.
* BFS on the full node graph is too slow for the constraints.
* The value-compression plus binary-lifting approach scales to `10^5` queries.
